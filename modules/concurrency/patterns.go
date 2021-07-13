package concurrency

import (
	"fmt"
	"time"
)

/*
	go routines allow concurrency. concurrency allows for parrellism.
	a go routine requires 2kb of memory, can grow and shrink accourding to need, can spawn hundreds to thousands per CPU Thread.
	go routine is cheap to create and destroy during runtime.

*/

func fibonacciRanged(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// the sender should close, never close at the receiver
	close(c) // channels don't need to be closed, are not necessary in all situations
	// it's necessary to close a connection when using a range for it to end successfully
}

func ExampleRangePattern() {
	c := make(chan int, 10)
	go fibonacciRanged(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func ExampleSelectPattern() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

func ExampleSimpleTimer() {
	timer := time.NewTimer(3 * time.Second)
	go func() {
		// timers setup a single delay to trigger
		// it it accomplished ONCE in the future
		<-timer.C
		fmt.Println("Timer fired")
	}()
	time.Sleep(2 * time.Second)
	// you can c
	stopped := timer.Stop()
	if stopped {
		fmt.Println("timer stopped")
	}
	fmt.Println("All done")
}
func ExampleSimpleTicker() {
	// tickers work on intervals
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for { // for loop to constantly accept ticker channel
			select {
			case <-done:
				return
			case tick := <-ticker.C:
				fmt.Println("Ticker fired at", tick)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("All done")
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "working on,", j)
		results <- j + 1
	}
}
func ExWorkerPool() {
	jobs := make(chan int)
	results := make(chan int)
	const numJob = 5

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	go func() {
		for j := 1; j <= numJob; j++ {
			jobs <- j
		}
		close(jobs) // we can close jobs since the for loop initiates the sending
		// never close on receiving.
	}()

	for a := 1; a <= numJob; a++ {
		//fmt.Println("Results", res)
		<-results
	}

}

func ExampleUnbufferedChan123() {
	unbuffered := make(chan string)

	defer close(unbuffered)
	// need to use go routines for syncronized unbuffered channels
	go func(s string) {
		unbuffered <- s
	}("second")

	go func(s string) {
		unbuffered <- s
	}("first")

	fmt.Println(<-unbuffered)
	fmt.Println(<-unbuffered)
}

func runsFunc(runs int) func(ch chan int) {
	return func(ch chan int) {
		for i := 0; i < runs; i++ {
			fmt.Println("written ", i, "to ch")
			ch <- i
		}
		close(ch)
	}
}

func ExTimeOuts() {
	// Do you need to buffer channel to prevent lockout if timedout?
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
		close(c1)
	}()

	go func() {
		select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1")
		}
	}()

	done := make(chan bool)
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
		close(c2)
	}()

	go func() {
		for {
			select {
			case res := <-c2:
				fmt.Println(res)
				done <- true
			case <-time.After(3 * time.Second):
				fmt.Println("timeout 2")
				done <- true
			default:
				fmt.Println("default")
			}
		}
	}()

	<-done
	close(done)
}

// Issue with multiple go routines might need to be enclosed in 1 loop
//loop
// go 1
// go 2
// are working with eachother

func runner(runs int, name string) chan<- int {
	out := make(chan int)
	go func() {
		for i := 0; i < runs; i++ {
			fmt.Printf("%d runs on %s", i, name)
		}
	}()
	return out
}

func ExUnbufferedChan() {
	runIt := 3
	uch := make(chan int)
	instance := runsFunc(runIt * 3)
	go instance(uch)
	// intentionall delay to simulate processing
	time.Sleep(time.Second * 1)

	for value := range uch {
		fmt.Println("read ", value, "from ch")
		time.Sleep(time.Second * 1)
	}

}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	done := make(chan bool)
	mid := len(data) / 2
	var left []int
	go func() {
		left = mergeSort(data[:mid])
		done <- true
	}()
	right := mergeSort(data[mid:])
	<-done
	return merge(left, right)
}

func ExampleMergeSort() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("%v\n%v\n", data, mergeSort(data))
}
