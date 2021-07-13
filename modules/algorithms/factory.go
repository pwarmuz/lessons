package algorithms

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// Factory pattern is a creational design pattern that provides an interface for creating objects into a superclass by allowing subclasses to alter the type of objects that are created
// The abstract factory defines an interface for creating distinct products but leaves the creation to contrete factory classes.
// simple factory example
type person struct {
	Name string
	Age  int
}

func (p person) Greeting() {
	fmt.Printf("Hello, I'm %s", p.Name)
}

func NewPerson(name string, age int) *person {
	return &person{name, age}
}
func SimpleFactory() {
	p := NewPerson("Akira", 33)
	p.Greeting()

}

// Interface factory
type greeter interface {
	Greet()
}

// person now fulfills the greeter interface
func (p person) Greet() {
	fmt.Printf("Hello, I'm %s from an interface.", p.Name)
}

func NewGreeter(name string, age int) greeter {
	return &person{name, age}
}

func InterfaceFactory() {
	p := NewGreeter("Akira", 33)
	p.Greet()
}

// Factory functions return defaults of a product

// newPersonFactory abstracts the age attribute away from the closure
func newPersonFactory(age int) func(name string) person {
	return func(name string) person {
		return person{name, age}
	}
}

func ExampleFactoryFunctions() {
	newBaby := newPersonFactory(1)
	baby := newBaby("Sophia")

	newTeen := newPersonFactory(16)
	teen := newTeen("Lola")
	fmt.Println(baby, teen)
}

// By returning multiple interfaces you can have multiple factories returning different implementations. This can help generate mocks.

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

func newHTTPClient() doer {
	return &http.Client{}
}

type mockClient struct{}

func (*mockClient) Do(req *http.Request) (*http.Response, error) {
	res := httptest.NewRecorder()
	return res.Result(), nil
}

// newMockClient lets you test code without actually making external HTTP Calls
func newMockClient() doer {
	return &mockClient{}
}

// Factory Generators are factories of factories
// This is useful when constructing instances of different structs or interfaces that are not mutually exclusive.
// Also helpful for when you want multiple factories with different defaults.

type animal struct {
	breed string
	age   int
}

type animalHouse struct {
	name       string
	squareFeet int
}

type animalFactory struct {
	breed string
	house string
}

func (af animalFactory) newAnimal(age int) animal {
	return animal{af.breed, age}
}

func (af animalFactory) newHouse(squareFeet int) animalHouse {
	return animalHouse{af.house, squareFeet}
}

func ExampleAnimalFactories() {
	dogFactory := animalFactory{"water", "kennel"}
	dog := dogFactory.newAnimal(2)
	kennel := dogFactory.newHouse(8)

	catFactory := animalFactory{"siamese", "house"}
	cat := catFactory.newAnimal(1)
	house := catFactory.newHouse(1300)

	fmt.Println(dog.breed, kennel.squareFeet)
	fmt.Println(cat.breed, house.squareFeet)
}

// Do not overuse factories
// Factories should be used in cases where there is value from frequent reuse.
// it's reasonable to just use default zero values of structs in a struct literal when there is no reuse
// see https://pkg.go.dev/sync?utm_source=godoc#WaitGroup or https://golang.org/src/net/http/client.go#L56 do not use factories to initialize data
