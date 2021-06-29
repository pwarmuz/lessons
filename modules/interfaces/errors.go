package interfaces

import "fmt"

// Error is the type of a parse error; it satisfies the error interface.
type Error string

// Error this fulfills the error interface
func (e Error) Error() string {
	return fmt.Sprint("customized error for: ", string(e))
}

// AltError is created as a type to force a panic during recovery at the type assertion
// change the type assertion e.(Error) to e.(AltError) to trigger
type AltError string

func (e AltError) Error() string {
	return fmt.Sprint("never gets called...", string(e))
}

// ForError just does something to create a panic
type ForError struct {
	e Error
}

func (fr *ForError) doSomething() (*ForError, error) {
	fmt.Println("executing panic")
	panic(fr.e)
	return &ForError{"This never returns"}, nil // this never returns because everything is reset in the recovery logic
}

func perform() (fr *ForError, err error) {
	fmt.Println("Creating custom error")
	fr = &ForError{"custom faked error"}
	// doParse will panic if there is a parse error.
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recovered from panic but could still fail the assertion")
			fr = nil        // Clear return value.
			err = e.(Error) // Will re-panic if not a parse error.
			// if you switch type Error to AltError it will panic with this response
			// panic: customized error for: custom faked error [recovered]
			// panic: interface conversion: interface {} is interfaces.Error, not interfaces.AltError
		}
	}()
	return fr.doSomething()
}

// Compile returns a parsed representation of the regular expression.
func ExampleErrorPanicRecovery() {
	fr, err := perform()
	fmt.Println("Completed for err", fr, "error", err)
}
