package loops

import "fmt"

func ExampleLoops() {
	var function func()
	var inside int
	for i := 0; i < 10; i++ {
		// because i is passed by reference
		// the value will be updated and not re-declared as i
		// however i is also being assigned to inside
		// inside will therefore be passed by value
		inside = i
		fmt.Println("within for loop", i, "with", inside, "inside")
		function = func() { fmt.Println("within anon func", i, "with", inside, "inside") }
	}
	// notice function() will result i value of 10
	// because i is incremented to 10 to stop the for loop (i++ is a post increment, go does not support pre increments)
	// however the inside value remains 9 because it assigned within the loop
	function()
	// calling it a second time to drive the point across
	function()
	// this effect can also happen with range loops as well
}

func ExampleLoopsPassedByValue() {
	f := func(i int) { fmt.Println("within anon func passed by value", i) }
	fByReference := func(i *int) { fmt.Println("within anon func passed by reference", *i) }
	var inside int
	var insideReference *int
	for i := 0; i < 10; i++ {
		inside = i
		f(i)
		insideReference = &i
		fByReference(&i)
	}
	// similar in concept to the above but i was passed by value into the function
	f(inside)

	// passed by reference
	fByReference(insideReference)
}
