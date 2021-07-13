package algorithms

import (
	"fmt"
	"sync"
)

// singleton pattern isn't as useful in Go because packages should be used to control single instance of variables and objects
// this implementation does that but implements it in an interesting way because no object is actually created because the struct is empty
// we are exploiting the methods of our singleton instance to mutate data and retrieve it.
// Notice no pointer is needed in the method receivers because global variables are being mutated
// This example also demonstrates why traditional singleton implementations are not as relevant in Go
// You could simply have a global variable in the package with a Get and Set function to manipulate it.
type singleton struct{}

var (
	instance *singleton
	once     sync.Once
	counter  int
	project  string
)

func (s singleton) AddCounter(i int) int {
	counter += i
	return counter
}

func (s singleton) GetCounter() int {
	return counter
}
func (s singleton) GetName() string {
	return project
}

func NewInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
		counter = 1
		project = "Mine"
	})
	return instance
}

func ExampleSingletons() {
	s := NewInstance()
	fmt.Println(s.AddCounter(4))
	s2 := NewInstance()
	fmt.Println(s2.GetCounter())
	fmt.Println(s2.GetName())
}
