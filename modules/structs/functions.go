package structs

import "fmt"

func ExampleStructImplementations() {
	// these implementation makes sense if you don't plan on using a struct for anything else
	// and you don't want confusion with a potentially similar struct that could be amended
	// these are anonymous struct types

	// zero value anonymous type option
	var oneInstance struct {
		int
		mutable int
		string
	}
	fmt.Println(oneInstance) // {0 0 }
	// can't access the unnamed int field and unamed string field
	oneInstance.mutable = 1  // this is the only field for manipulation
	fmt.Println(oneInstance) // {0 1 }
	oneInstance.mutable = 2
	copy := oneInstance
	oneInstance.mutable = 3
	fmt.Println(copy.mutable) // prints 2
	// a ridiculous way to change the values
	oneInstance = struct {
		int
		mutable int
		string
	}{
		3,
		4,
		"ridiculous",
	}
	fmt.Println(oneInstance) // {3 4 ridiculous}

	// If you want non-zero valued data then it needs to be constructed
	constructedValues := struct {
		int
		mutable string
	}{
		1,
		"something",
	}
	fmt.Println(constructedValues) // {1 something}
	constructedValues.mutable = "changed"
	fmt.Println(constructedValues) // {1 changed}

}

// this is a signature of a user defined function type
type action func(initial int) (result int, err error)

//A function can use other functions as arguments and return values.
// this is a higher-order function signature
type strategy func(int) action

type functionLiteralArg func(less func(i, j int) bool) // less is a function literal argument

func ExampleHighOrderFunc() {
	// stay fulfills the action type
	modifier := 2
	stay := func(i int) (int, error) { // this function literal uses a modifier from outside it's declaration and assignment but within scope
		return i * modifier, nil // anytime stay is used the modifier will be used
	}
	// pointless higher-order function
	highStay := func() action {
		return stay
	}

	var highRoll strategy
	highRoll = func(start int) action { // fulfills the strategy signature and is a higher-order function
		i := start * 10                         // function literals are closures, they may use variables outside the function; in this case i is being used
		return func(initial int) (int, error) { // anonymous function and function literal
			i += initial // i is being used to sum up initial which is part of the closure
			return i, nil
		}
	}

	var value int
	var err error
	stayAction := highStay()   // high stay simply returns an action
	value, err = stayAction(1) // the action type as a function takes 1 as an initial value to get a result
	if err != nil {
	} // pointless err just to demonstrate that action returned 2 values and is being handled
	fmt.Println("stay value", value) // value is 2

	modifier = 6
	value, err = stayAction(1)
	fmt.Println("stay value modified", value) // value is 6 because modifier was changed

	rollAction := highRoll(1)
	value, err = rollAction(10)
	if err != nil {
	}
	fmt.Println("roll value", value) // value is 20; initial 1*10=10 from rollResult is counted towards the action type function rollResult which adds 10 to that

}
