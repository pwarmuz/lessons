package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"

	"www-asteria/app"

	"github.com/caddyserver/certmagic"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// HTTPHandle unused, intended for multi-domain
type HTTPHandle struct {
	Handler http.Handler
}

// Muxer object
type Muxer struct {
	Router    *httprouter.Router
	TLSConfig *tls.Config
	Server    *http.Server
	FS        http.FileSystem
	ACME      *certmagic.ACMEManager
	HTTP      *http.Server
	Suga      *zap.SugaredLogger
	Hosts
	sync.Mutex
}

// TPF template parsed files
var TPF *template.Template

// PPF public parsed files
var PPF *template.Template

// init
// Read files once
// Build memory map
// Reuse memory map
func init() {
	suga := app.ZapLog()
	suga.Info("INIT", "Mux Initiated, Suga logging")
}

// main creates the Mux and Server
// Initializes configuration
// Executes listeners
// HTTP listener is created within a concurrent instance and redirects HTTP -> HTTPS
// and exempts the HTTP ACME challenge from hard sanitation into the domain name
//
// HTTPS listener is created at the end of main and inheriting all configuration
func main() {
	// Load the configuration file
	suga := app.ZapLog()
	suga.Infow("Web Server Starting in... ", "Environment", app.ENVIRONMENT.Release())
	suga.Info("Reload the webserver with Ctrl+C and re-init to see HTML changes")
	defer suga.Info("...Web Server Stopped")

	MUX := NewMux()

	// Production or Localhost selection performed here
	// Production or Localhost Server configuration is the result
	MUX.configurationSelection()

	// Asset allocation
	MUX.assetAllocation()
	// Server Configuration and Initialization
	MUX.ServerInit()

	// sigint interruption handling
	idleHTTPSClosed := make(chan struct{})
	// Shutdown Connection handling go routine
	// Contains sigint trigger for shutdown acceptance
	go func() {
		// sigint interruption
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		// sigint received time to gracefully shutdown
		fmt.Println("") // This cleans up the log by ending the sigint character in console
		suga.Infow("Closing...", "Environment", app.ENVIRONMENT.Release())
		MUX.shutdownHTTPS()
		MUX.shutdownHTTP()
		close(idleHTTPSClosed)
	}()

	// Independent instance for HTTP Redirect
	// Shutdown is called within sigint go routine
	go func() {
		MUX.httpInstanceSelection()
	}()

	// HTTPS Server ListenAndServeTLS
	if err := MUX.Server.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
		// Error starting or closing listener:
		suga.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleHTTPSClosed
	suga.Info("HTTPS/HTTP CONNECTIONS CLOSED...")
}

// ServeHTTP is the highest level ServeHTTP for the MUX and is part of the http.Handler interface
// This gets called at a high level which gets passed into asset serving and other ServeHTTP requests
// Since this is high level and affects ServeFiles (to be tested) will this prevent other domains from linking assets?
//
// Reverse Proxy info regarding ACME Challenges
// You must ensure that incoming validation requests containt the correct value for the HTTP Host header.
// If you operate lego behind a non-transparent reverse proxy (such as Apache or NGINX),
// you might need to alter the header field using --http.proxy-header X-Forwarded-Host. [return]
func (mx *Muxer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// PARSING URL
	//	fmt.Println(url.QueryUnescape(`melissajwolff+testabc@gmail.com`))
	// melissajwolff testabc@gmail.com <nil>
	// fmt.Println(url.ParseQuery(`auth-token=KeQEtE8VI3duSqXHnrDoFKMeilRfmE&auth-email=melissajwolff+testabc@gmail.com`))
	// map[auth-email:[melissajwolff testabc@gmail.com] auth-token:[KeQEtE8VI3duSqXHnrDoFKMeilRfmE]] <nil>
	// fmt.Println(url.QueryEscape(`melissajwolff+testabc@gmail.com`))
	// melissajwolff%2Btestabc%40gmail.com
	//
	// HTTP OPTIONS
	// curl -i --request OPTIONS https://localhost:10443/submissionCollaborate responds with allow: OPTIONS, POST
	//
	// MUTEX
	// deferring the mutex unlock lets handler := h.list[host] be handled within the if block
	// additionally allows for handing any := assignment within the if block
	mx.Lock()
	defer mx.Unlock()
	host := req.Host

	if ok := mx.whitelist[host]; ok {
		// methodOverride, you can create other functions with a bool
		// methodOverride(w, req, ps, bool) that toggles it via ADMIN control
		methodOverride(w, req)

		if !strings.HasPrefix(req.URL.String(), "/static/") {
			mx.Suga.Debugw("MUX ServeHTTP", "Type", "Whitelist", "Host", host, " Requested", req.URL.String())
		}
		// URL information should be sanitized and formatted properly here
		// NOTE, Using declared mutexes doesn't do shit
		if !app.ENVIRONMENT.Production() {

			if host == "philwarmuz.localhost:10443" {
				// Redirect should not be a literal copy of initial host path, http.Redirect will loop
				// E.g. blog.example.com should redirect to example.com/blog where the literal blog.example.com hostname won't loop with example.com
				//targetUrl := url.URL{ Scheme: "https", Host: req.Host, Path: r.URL.Path, RawQuery: r.URL.RawQuery, }
				//targetUrl := url.URL{Scheme: "https", Host: req.Host, Path: "/ethos"}
				http.Redirect(w, req, "https://localhost:10443/ethos", http.StatusFound)
			}
		}
		// Controls the domain, sets precidence over other ServeHTTP
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		w.Header().Set("Strict-Transport-Security", "max-age=15768000 ; includeSubDomains")
		mx.Router.ServeHTTP(w, req)
	} else {
		// Does this cause acme autocert to waste time finding certs for domains I don't want https handshaked?
		mx.Suga.Debugw("MUX ServeHTTP", "Type", "Forbidden", "Host", host, " Requested", req.URL.String())
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

// GZipServeFiles is a poorly attempted fileserver
// Muxer ServeHTTP handles the GZIP and HTTPS headers
// The fileserver inherits the gziphandler.GzipHandler() from ServerHandler = Muxer
// TODO: Note: MUX.ServeHTTP -> Router.GET(), MUX.ServeHTTP is Parent some filtering is best done there
// TODO: httpRouter ps Param cleanup
func (mx *Muxer) GZipServeFiles(path string, root string) {
	if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
		panic("path must end with /*filepath in path '" + path + "'")
	}
	// TODO: filepath mainly filters strings, Abs will get the path during compilation at development time
	fullPath := "FALSE IMPRESSION"
	mx.Suga.Debugw("Fileserver created", "URL Path", path, "Full Path as Root", fullPath)

	// http.Dir() converts string into http.FileSystem
	fileserver := http.FileServer(filterFS{http.Dir(root)})
	f := func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		req.URL.Path = ps.ByName("filepath")
		/*
			Tried to implement precompressed .gz files
			If gzip was accepted, then it would check for .gz files
			and serve them instead of the regular version
			Fails because it needs to decompress it and rename it to be properly interpreted as source file
			if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
				s := strings.Split(path, "/")
				mx.Suga.Debugw("Accepted Gzip |" + string(http.Dir(s[1])) + ps.ByName("filepath") + ".gz")
				if ok := tools.FileExists(string(http.Dir(s[1])) + ps.ByName("filepath") + ".gz"); ok {
					req.URL.Path = ps.ByName("filepath") + ".gz"
				}
			}
		*/

		// This file and directory validation enforces the appropriate
		// overwrite 404 Redirect page.
		fp := filepath.Join("static", filepath.Clean(req.URL.Path))
		info, err := os.Stat(fp)
		if err != nil {
			if os.IsNotExist(err) {
				http.ServeFile(w, req, "public/404.html")
				return
			}
		}
		if info.IsDir() {
			http.ServeFile(w, req, "public/404.html")
			return
		}

		// Prevent hotlinking
		// Ensure google and searchengines can browse and hotlink
		// Static folder redirect and Image redirect
		host := req.Host
		//Item HOST is >philwarmuz.com< Referrer >https://philwarmuz.com/static/css/custom.white.min.css
		// Need to strip /static.... ensure https:// is included
		// TODO: implement static url hashing, hash is generated every instance preventing hotlinking
		u, err := url.Parse(req.Header.Get(ReferrerReq))
		if err != nil {
			mx.Suga.Fatal(err)
		}

		mx.Suga.Debugw("GZip Serving File", "Path", req.URL.Path, "Host", u.Hostname(), "Scheme", u.Scheme)
		if ok := mx.referrer[u.Hostname()] || mx.whitelist[u.Hostname()] && u.Scheme == "https"; !ok {
			w.Header().Set("Connection", "close")
			//url := "https://localhost:10443/static/img/ahahah.jpg"
			url := "https://" + host + "/ahahah"
			//http.StatusTemporaryRedirect
			http.Redirect(w, req, url, http.StatusMovedPermanently)
			return
		}
		// Sanitized for serving
		fileserver.ServeHTTP(w, req)
	}
	// ps Params are dependent on the registered URL path
	// ps Params is using a catch-all *filepath which allows for /img/large/ to be found within /static/ where set where ps.ByName("filepath")
	// ps Params would not find /static/img/large if /static/:filepath where set where ps.ByName("filepath")
	mx.Router.GET("/static/*filepath", f)
}

func (mx *Muxer) shutdownHTTPS() {
	// We received an interrupt signal, shut down.
	if err := mx.Server.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		mx.Suga.Fatalf("HTTPS Server Shutdown: %v", err)
	}
	mx.Suga.Info("HTTPS Server Shutdown gracefully...")
}
func (mx *Muxer) shutdownHTTP() {
	// We received an interrupt signal, shut down.
	if err := mx.HTTP.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		mx.Suga.Fatalf("HTTP Server Shutdown: %v", err)
	}
	mx.Suga.Info("HTTP Server Shutdown gracefully...")
}

type filterFS struct {
	fs http.FileSystem
}

// func filterFS(root http.FileSystem) *filterFS {
// 	return &filterFS{root, ""}
// }
func (ffs filterFS) Open(path string) (http.File, error) {
	f, err := ffs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := ffs.fs.Open(index); err != nil {
			return nil, err
		}
		//return nil, err
	}
	return f, nil
}

type whitelist map[string]bool

// Hosts contains all host information for mux
type Hosts struct {
	whitelist
	referrer whitelist
}

// String for whitelists
// prints out the key name as a variadic list []string{}
func (e *whitelist) String() (temp string) {
	last := len(*e)
	cnt := 0
	for name := range *e {
		temp += name
		cnt++
		if cnt != last {
			temp += " "
		}
	}
	return
}
