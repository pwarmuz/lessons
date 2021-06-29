package algorithms

import (
	"fmt"
)

// Stack
// commonly refers to First In, Last Out (FILO) or Last In, First Out(LIFO)
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type Stack struct {
	slice []int
}

func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}
func (s *Stack) Pop() int {
	pop := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return pop
}
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

func (s Stack) String() string {
	return fmt.Sprint(s.slice)
}

func ExampleStack() {
	s := new(Stack)
	s.Push(100)
	s.Push(22)
	s.Push(37)
	s.Push(54)
	fmt.Println(s)
	fmt.Println(s.Peek())
	fmt.Println("Pop", s.Pop())
	fmt.Println(s)
	fmt.Println("Pop", s.Pop())
	fmt.Println(s.Peek())
	fmt.Println(s)
}
