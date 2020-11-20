package algorithms

import "fmt"

type queue []int

func (q *queue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *queue) Dequeue() int {
	// (*q)[n] means [n] is converted to *Point
	dequeue := (*q)[0]
	*q = (*q)[1:len(*q)]
	return dequeue
}
func (q *queue) String() string {
	return fmt.Sprint(*q)
}

// This looks better than the above implementation

type altQueue []int
type AltQueue struct {
	altQueue
}

func (q *AltQueue) Enqueue(i int) {
	q.altQueue = append(q.altQueue, i)
}
func (q *AltQueue) Dequeue() int {
	dequeue := q.altQueue[0]
	q.altQueue = q.altQueue[1:len(q.altQueue)]
	return dequeue
}
func (q AltQueue) String() string {
	return fmt.Sprint(q.altQueue)
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
