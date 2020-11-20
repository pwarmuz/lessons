package defering

import "fmt"

func ExampleDefering() {
	// notice the extra ()
	// because setupTearDown is a closure
	// the returned function is what is meant to be deferred
	defer setupTearDown()()
	// this is because defer takes a function call

	// the above is equivalent to
	// f := setupTearDown()
	// defer f()
}

func setupTearDown() func() {
	fmt.Println("something setup")
	return func() {
		fmt.Println("something torn down")
	}
}
