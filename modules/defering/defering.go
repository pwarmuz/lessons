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

// trace will be called despite it being
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

// last will only be called when it defers
func last(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer fmt.Println("deferred a") // notice no function arguments, this is called at defer
	fmt.Println("in a")
}

func b() {
	fmt.Println("Starting b")
	defer last(trace("b")) // Notice the last function and trace function, trace will be called after the previous line
	fmt.Println("in b")
	a()
}

func ExampleDeferTracing() {
	b()
}
