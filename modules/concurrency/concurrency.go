package concurrency

// Go has support for concurrent programming

// In the previous example we looked at setting up a simple
// [HTTP server](http-servers). HTTP servers are useful for
// demonstrating the usage of `context.Context` for
// controlling cancellation. A `Context` carries deadlines,
// cancellation signals, and other request-scoped values
// across API boundaries and goroutines.
import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// A `context.Context` is created for each request by
	// the `net/http` machinery, and is available with
	// the `Context()` method.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	fmt.Fprintf(w, "initial load\n")
	defer fmt.Println("server: hello handler ended")

	// Wait for a few seconds before sending a reply to the
	// client. This could simulate some work the server is
	// doing. While working, keep an eye on the context's
	// `Done()` channel for a signal that we should cancel
	// the work and return as soon as possible.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "context case\n")
	case <-ctx.Done():
		// The context's `Err()` method returns an error
		// that explains why the `Done()` channel was
		// closed.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func ExampleCurrency() {
	// As before, we register our handler on the "/hello"
	// route, and start serving on the default mux
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

// Buffered and Unbuffered channels
// by default channels are unbuffered
func ExampleUnbufferedChan() {
	// notice that this channel does not have a capacity
	// this declares it an unbuffered channel
	unbuffered := make(chan string)
	// second independent go routine will actually go second
	// despite being created first
	go func(s string) {
		unbuffered <- s
		// close(unbuffered)
		// this above, close(unbuffered) is dangerous IF
		// this go routine truely didn't go second you would be closeing before all information was processed
		// it's only safe BECAUSE it does go second
		// you can uncomment it, comment the defer closure below and it will work but it's not a smart place to put
	}("second")
	// this go routine will be processed first
	go func(s string) {
		unbuffered <- s
	}("first")
	// since close should be executed by the sender a smarter thing to do would be to defer close it after both go routines were created
	defer close(unbuffered)
	// each one of these individual println's are required to read each instance of the data
	fmt.Println(<-unbuffered)
	fmt.Println(<-unbuffered)
}

// Buffered and Unbuffered channels
// by default channels are unbuffered
func ExampleBufferedChan() {
	// notice that this channel does not have a capacity
	// this declares it an unbuffered channel
	unbuffered := make(chan string, 2)
	// since "unbuffered" channels are used within this function
	// and close should be executed by the sender
	// the sender is actually towards the end as "unbuffered <- "eight""
	// so defering this here is valid
	defer close(unbuffered)
	// this is buffered so it can be instantianted within the current function
	unbuffered <- "first"
	unbuffered <- "second"
	// unbuffered <- "third"  //this won't work here because their is no corresponding reciever, remember 2 channels were declared
	// this "third" value will create a deadlock

	// each one of these individual println's are required to read each instance of the data
	fmt.Println(<-unbuffered)
	fmt.Println(<-unbuffered)

	// these additional channels will work because 2 channels have already been recieved
	unbuffered <- "third"
	unbuffered <- "forth"
	fmt.Println(<-unbuffered)
	fmt.Println(<-unbuffered)

	// this equally works because a reciever is immediatley found after
	unbuffered <- "fifth"
	fmt.Println(<-unbuffered)
	unbuffered <- "sixth"
	fmt.Println(<-unbuffered)
	unbuffered <- "seventh"
	fmt.Println(<-unbuffered)
	unbuffered <- "eight" // this is the last sender channel
	fmt.Println(<-unbuffered)
}

func runs(runs int) func(ch chan int) {
	return func(ch chan int) {
		for i := 0; i < runs; i++ {
			ch <- i
			fmt.Println("written ", i, "to ch")
		}
		close(ch)
	}
}
func ExampleBufferedChanRoutine() {
	capacity := 3
	buffered := make(chan int, capacity)
	instance := runs(capacity * 3)
	go instance(buffered)
	// intentionall delay to simulate processing
	time.Sleep(time.Second * 1)

	// read values
	for value := range buffered {
		fmt.Println("read ", value, "from ch")
		time.Sleep(time.Second * 1)
	}
	// output will vary
	// generally the trend will start with 3 channels written intially, where 3 is the capacity
	// generally the trend will end with 3 channels read, where 3 is the capacity
	// there will be chances of 2 items written one after the other and 2 items read one after the other
}
