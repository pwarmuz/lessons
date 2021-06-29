package conversions

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// All conversions between typed expressions must be explicitly stated
// and illegal conversions are caught by the compiler

// ExampleConversionCosts
func ExampleConversionCosts() {
	// Conversions to and from numbers and strings may change the representation
	// but will have some run-time cost.
	// Other conversions will change the type but not the representation.
	start := time.Now()
	convert(cFloat32, 14)
	convert(cString, 1)
	end := time.Now()
	fmt.Println("Process took", end.Sub(start))
}

func cFloat32(i int) string {
	fmt.Println("float", float32(i))
	return strconv.Itoa(i)
}
func cString(i int) string {
	fmt.Println("string", string(byte(i)))
	return string(byte(i))
}

func convert(funky func(int) string, mult int) {
	for i := 0; i == 100; i++ {
		fmt.Println(funky(i * mult))
	}
}

// the only implicit conversion is from an untyped constant into a required type
const exampleImplicit = 1

// Go language is strongly typed which means you need declare a type implicitly or explicitly
// It's strict when it comes to types and will report errors during compilation.

type customString string

// Constants can be untyped and can be manipulated by assuming a type
// don't get to clever with this because assigning an untyped float64 constant into a float32 variable for example could cause problems
const (
	cstring               = "my string"
	cint                  = 1
	cstring2 customString = "hello"
)

func ExampleAliasing() {
	// golang uses
	// byte as an alias for uint8
	// rune as an alias for int32 to represent a unicode code point
	var b byte
	var ui8 uint8 = 1
	b = ui8
	fmt.Println("byte:", reflect.TypeOf(b).Name(), "|| uint8:", reflect.TypeOf(ui8).Name())

	type myInt32 int32
	type MyAliasI32 = int32
	var myInt myInt32 = 105
	var myAlias MyAliasI32 = 32
	// myInt = myAlias cannot assign type int32 to type myInt32
	myInt = myInt32(myAlias) // assigned and converted
	fmt.Println(myInt)
	var i32 int32 = 100
	i32 = myAlias // no conversion need because alias is the same type
	fmt.Println(i32)
}

func ExamplePointerConversion() {
	type MyInt int
	type IntPtr *int
	type MyIntPtr *MyInt

	var pi = new(int)
	var ip IntPtr = pi
	// cannot implicitly convert as var _ *MyInt = pi
	var _ = (*MyInt)(pi) // must be done explicitly
	var _ IntPtr = pi    // is the same type

	// cannot convert *int to MyIntPtr directly
	// var _ MyIntPtr = pi  // won't convert implicitly
	// var _ = MyIntPtr(pi) // won't covert explicitly
	// if pi is converted to *MyInt then it is equal to MyIntPtr
	// when done indirectly as seen with indirection ( )
	var _ MyIntPtr = (*MyInt)(pi)  // A) both A & B work equally
	var _ = MyIntPtr((*MyInt)(pi)) // B) equal as above

	// values of IntPtr cannot be converted to MyIntPtr directly
	// where ip was IntPtr
	// var _ MyIntPtr = ip
	// var _ = MyIntPtr(ip)
	// but IntPtr can be converted indirectly
	var _ MyIntPtr = (*MyInt)((*int)(ip))  // indirectly achieved
	var _ = MyIntPtr((*MyInt)((*int)(ip))) // indirectly achieved
}

func ExampleConversions() {
	// If an underlying type is the same then conversion is valid
	type underlying string
	var us string = "hello"
	under := underlying(us)
	fmt.Println("underlying", reflect.TypeOf(under).Name())

	// however
	// us = underlying("string")
	// will not work because us is already declared a string and not of type underlying
	// another example would be map[string]int =/= map[mystringtype]int where mystringtype of type string
	// you would need to convert mystringtype to a string before using as a key for the map

	// you can express the conversion explicitly
	var f = func(n [2]int) {
		fmt.Println(n)
	}

	type T [2]int
	var v T
	// where, notice func(n [2]int) is not accepting T
	f(v)
	// is the same as
	f([2]int(v))

	var _ = func() {
		type T *int64
		var n int64 = 1
		var m int64 = 2
		var p T = &n
		var q *int64 = &m
		p = q
		fmt.Println(*p)

		type X map[string]int
		var x X
		var y map[string]int
		// since X is an underlying map[string]int you can assign y to it as only one is named
		x = y
		_ = x
		// however
		type Z map[string]int
		//var z Z
		//x = z  caanot use type Z since since it doesn't match X as both are named
	}

	type AnotherString string // local scope type
	var mn customString
	mn = cstring

	var myStringy string
	myStringy = string(cstring2)

	var mfloat float64
	mfloat = cint

	fmt.Printf("custom string, float, converted string %s, %f %s", mn, mfloat, myStringy)
	fmt.Println("reflect", reflect.TypeOf(mn).Name(), reflect.TypeOf(mfloat).Name())
	var another AnotherString
	another = cstring
	fmt.Println("another", reflect.TypeOf(another).Name(), reflect.TypeOf(mfloat).Name())

	//	var n int = 1
	//	var p *int
	//	(*int(n))
	// t := 2 * time.Second is the same as t := time.Duration(2) * time.Second since types cannot be mixed and in this case can only multiple time.Duration

}
