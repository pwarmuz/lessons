package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func access(ch chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println("start accessing channel")

	for i := range ch {
		wg.Done()
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func ExampleWGLoop() {
	// un-buffered is to maintain synchronization
	//ch := make(chan int)
	// buffered is good for queueing
	//ch := make(chan int, 3)
	ch := make(chan int, 3)

	//defer
	var wg sync.WaitGroup
	go access(ch, &wg)

	for i := 0; i < 9; i++ {
		// time.Sleep(time.Second) both buffered and unbuffered will block the channel until it is filled
		wg.Add(1)
		ch <- i
		fmt.Println("Filled")
	}
	close(ch)
	wg.Wait()

	fmt.Println("ender")
}
