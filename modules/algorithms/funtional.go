package algorithms

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// functional options patterns in Go
// They take the form of extra arguments to a function that extend it's behavior.
// e.g. h:=NewHouse(WithPool(),WithGranite())
// Benefits:
// 	being explicit with function name for construction without exposing internal data
// 	allowing extensible functionality with additional arguments in option function e.g. hardwoodRoom(size int) houseOptions
//		flexibility because it doesn't conform to a specific order of arguments when constructing
//		beneficial pattern for API

type material struct {
	name string
	cost int
	size int
}
type house struct {
	kitchen    material
	pool       material
	livingRoom material
}

type houseOptions func(*house)

func newHouse(opts ...houseOptions) *house {
	var (
		k = material{name: "steel", cost: 1000, size: 300}
		p = material{name: "plaster", cost: 1000, size: 300}
		l = material{name: "carpet", cost: 1000, size: 300}
	)
	h := &house{kitchen: k, pool: p, livingRoom: l}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func graniteKitchen() houseOptions {
	return func(h *house) {
		h.kitchen.name = "granite"
	}
}
func tilePool() houseOptions {
	return func(h *house) {
		h.kitchen.name = "tile"
	}
}
func hardwoodRoom(size int) houseOptions {
	return func(h *house) {
		h.livingRoom.name = "hard wood"
		h.livingRoom.size = size
	}
}

func SimpleFunctionalPattern() {
	h := newHouse(graniteKitchen(), tilePool(), hardwoodRoom(250))
	fmt.Println(h.kitchen.name, h.livingRoom.name, h.pool.name)
}

// NewServer functional option pattern

type srvOpts func(*Server)

type Server struct {
	listener net.Listener
	client   http.Client
}

func timeout(t int) srvOpts {
	return func(srv *Server) {
		srv.client.Transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   time.Duration(t) * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
	}
}

func newServer(addr string, options ...srvOpts) (*Server, error) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	srv := &Server{listener: listen}

	for _, option := range options {
		option(srv)
	}
	return srv, nil
}

// to := timeout(30)
// srv, err := newServer("80", to)
