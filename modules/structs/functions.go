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
