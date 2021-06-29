package interfaces

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Exportable struct {
	PubInt        int
	PubString     string
	privateInt    int
	privateString string
}

// consider the functionality(actions) that is common between datatypes
// interface{} is not any type but a type of interface{}
// interfaces are two words wide (type, value)
// returning an interface{} is generally not practical
// accepting an interface{}, like marshalJSON(interface{}), is more practical
//
// Go does not use inheritance like C++ however it uses composition to create objects into complex ones.
// Interfaces in Go provide a method of organizing complex compositions, which allow common, reusable code.
// https://simplyian.com/2016/06/28/Why-Go-s-structs-are-superior-to-class-based-inheritance/
// https://odetocode.com/blogs/scott/archive/2019/01/03/composition-over-inheritance-in-go.aspx
// Mocked up go playground https://play.golang.org/p/aptoqZJhy9L
//
// Interfaces should be small
// Interfaces should strictly define what that interface is; not to include defintions of what the interface is not
// e.g. an interface of animal should not have a method isThisADuck()
// Interfaces are not classes, they should be slimmer, they don't have constructors or destructors
// Interfaces define function signatures and not the underlying behavior

// Processor
type Processor interface {
	Process() string // "Process() string" method must be implemented to satisfy Processor interface
}

// DoIt uses Processor as an argument to print information
func DoIt(p Processor) {
	fmt.Printf(p.Process())
}

type Remote struct {
	values int
}

func (r *Remote) add(i int) {
	r.values = i
	DoIt(r)
}

// Process satisfies the Processor interface
func (r *Remote) Process() string {
	return fmt.Sprintf("This remote is %d\n", r.values)
}

type Local struct {
	values int
}

func (l Local) add(i int) {
	l.values = i
	DoIt(l)
}

// Process satisfies the Processor interface
func (l Local) Process() string {
	return fmt.Sprintf("This local is %d\n", l.values)
}

type Controller struct {
	values    int
	Processor // Processor is embedded into Controller struct, notice it is not named
}

func (c *Controller) add(i int) {
	c.values = i
	DoIt(c)
}

func (c Controller) Process() string {
	return fmt.Sprintf("This Controller is %d\n", c.values)
}

func ExampleInterface() {
	var remote Remote
	remote.add(1)
	var local Local
	local.add(2)
	DoIt(&remote) // remote uses a pointer in it's receiver so we need to de-reference
	DoIt(local)   // local is passed by value in it's receiver

	//Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	// var nilRemote *Remote
	// nilRemote.add(1)  <- will not compile panic: runtime error: invalid memory address or nil pointer dereference
	// pointers, interfaces, channels, functions, maps, and slices types can be nil
	// map, slice and function types don’t support comparison
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		switch {
		case p[i] >= 'A' && p[i] <= 'Z':
			p[i] = 'A' + (p[i]-'A'+13)%26
		case p[i] >= 'a' && p[i] <= 'z':
			p[i] = 'a' + (p[i]-'a'+13)%26
		}
	}
	return
}

func ExampleRot13() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// Rot13Converter is used instead of making a struct and implementing an interface since rot13Reader's only purpose is to make a rot13 conversion
// and only implements Read from the Reader interface
func Rot13Conversion(in io.Reader, b []byte) (string, error) {
	_, err := in.Read(b)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for i := 0; i < len(b); i++ {
		switch {
		case b[i] >= 'A' && b[i] <= 'Z':
			b[i] = 'A' + (b[i]-'A'+13)%26
		case b[i] >= 'a' && b[i] <= 'z':
			b[i] = 'a' + (b[i]-'a'+13)%26
		}
	}
	return string(b), nil
}

func ExampleRot13NoStruct() {
	code := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(code)
	b := make([]byte, len(code))
	rot13, err := Rot13Conversion(s, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("converted bytes", string(b), "or returned string", rot13)
}

type MyEmptyInterface interface{}

type MyAltStringer interface {
	Stringed() string
}

type MyStringType string

func (s MyStringType) Stringed() string { // implements MyAltStringer
	return fmt.Sprint("Stringed from MyAltStringer ", string(s))
}
func (s MyStringType) String() string { // implements fmt.Stringer
	return fmt.Sprint("Stringed from fmt.Stringer ", string(s))
}

func ExampleAssertion() {
	// creating an empty interface and associating it to a string
	// this part shows the interface is what is being type asserted
	// it shows the result and validity options as returns
	var emptyInterfaceString MyEmptyInterface = "type interface as string"
	r, ok := emptyInterfaceString.(string) // direct type assertion, we know checkThat is a string
	fmt.Println(r, ok)                     // type interface as string true

	// this part shows that an interface parameter will accept it as an interface and allow for type assertion
	// a type switch is most appropriate when accepting an interface as you generally would accept multiple ones
	// if you weren't intending to accept multiple types it would not be smart to use an interface as a parameter
	var checkThis string = "type is string"
	// its important to understand that the interface is be asserted
	checkFunc := func(i interface{}) {
		r, ok := i.(string) // we know that the interface will be a string
		fmt.Println(r, ok)  // type is string true
	}
	checkFunc(checkThis) // this is just an example of a string type being accepted as an interface and being asserted as a string within it

	// this part shows how various interfaces are fulfilled
	var typeStringWithMethods MyStringType = "MyStringType with methods" //this becomes a member of MyAltStringer and fmt.Stringer because it fulfils those methods

	returnedMyStringType := func(i interface{}) MyStringType { // remember, an interface is not "anything" it's of type interface
		r, ok := i.(MyStringType)
		if !ok {
			return MyStringType("Not of MyStringType")
		}
		return r
	}
	returnedMyAltStringer := func(i interface{}) MyAltStringer {
		r, ok := i.(MyAltStringer)
		if !ok {
			return MyStringType("Not of MyAltStringer interface")
		}
		return r
	}
	returnedFmtStringer := func(i interface{}) fmt.Stringer {
		r, ok := i.(fmt.Stringer)
		if !ok {
			return MyStringType("Not of fmt.Stringer interface")
		}
		return r
	}

	returnedAsStringConversion := func(i interface{}) string {
		convert := i.(MyStringType) // is asserted as MyStringType but becomes converted
		// ok omitted because this is controlled example
		return string(convert)
	}

	/*
		tryPointer := func(i *interface{}) bool {
			return true
		}
		oops := &typeStringWithMethods
		f := tryPointer(oops)  // oops won't work, can't be used with *interface{}
		fmt.Println("got tried", f)
	*/

	a := returnedMyStringType(typeStringWithMethods)
	fmt.Println(a.Stringed()) // Stringed from MyAltStringer MyStringType with methods

	b := returnedMyAltStringer(typeStringWithMethods)
	fmt.Println(b.Stringed()) // Stringed from MyAltStringer MyStringType with methods

	c := returnedFmtStringer(typeStringWithMethods)
	fmt.Println(c) // notice not .Stringed() or .String() because the String method is fulfilling the fmt.Stringer interface and Println executes it.
	// Stringed from fmt.Stringer MyStringType with methods
	d := returnedFmtStringer(emptyInterfaceString)
	fmt.Println(d) // this fails the assertion but is converted to MyStringType to fulfill the return
	// Stringed from fmt.Stringer Not of fmt.Stringer interface
	e := returnedAsStringConversion(typeStringWithMethods)
	fmt.Println(e, "as a string conversion with default .String() implementation") // as a string is fulfills it's default .String() implementation and the variable is displayed
	// MyStringType with methods as a string conversion with default .String() implementation

}
