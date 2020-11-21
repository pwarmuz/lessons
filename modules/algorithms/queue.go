package algorithms

import "fmt"

// queue
// commonly referred to as First In, First Out (FIFO)
// https://en.wikipedia.org/wiki/FIFO_(computing_and_electronics)
type queue []int

// Enqueue adds the item to the back of the order
func (q *queue) Enqueue(i int) {
	*q = append(*q, i)
}

// Dequeue removes the item from the front of the order
func (q *queue) Dequeue() int {
	// (*q)[n] means [n] is converted to *Point
	dequeue := (*q)[0]
	*q = (*q)[1:len(*q)]
	return dequeue
}

func (q *queue) String() string {
	return fmt.Sprint(*q)
}

// AltQueue
// This is a alternative implementation pattern
type AltQueue struct {
	intQueue
}
type intQueue []int

func (q *AltQueue) Enqueue(i int) {
	q.intQueue = append(q.intQueue, i)
}

func (q *AltQueue) Dequeue() int {
	dequeue := q.intQueue[0]
	q.intQueue = q.intQueue[1:len(q.intQueue)]
	return dequeue
}

func (q AltQueue) String() string {
	return fmt.Sprint(q.intQueue)
}

func ExampleQueue() {
	var q *queue = new(queue)
	fmt.Println("Queue is first in, first out")
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(9)
	fmt.Println(*q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(*q)

	var q2 queue
	fmt.Println("Queue is first in, first out, already ptr")
	q2.Enqueue(5)
	q2.Enqueue(6)
	q2.Enqueue(9)
	fmt.Println(q2)
	fmt.Println(q2.Dequeue())
	fmt.Println(q2.Dequeue())
	fmt.Println(q2.Dequeue())
	fmt.Println(q2)

}
func ExampleAltQueue() {
	var q *AltQueue = new(AltQueue)
	fmt.Println("Queue is first in, first out111")
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(9)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q)

	var q2 AltQueue
	fmt.Println("Queue is first in, first out 222")
	q2.Enqueue(5)
	q2.Enqueue(6)
	q2.Enqueue(9)
	fmt.Println(q2)
	fmt.Println(q2.Dequeue())
	fmt.Println(q2.Dequeue())
	fmt.Println(q2.Dequeue())
	fmt.Println(q2)
}
