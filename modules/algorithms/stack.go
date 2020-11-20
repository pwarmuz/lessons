package algorithms

import (
	"fmt"
	"reflect"
)

type Stack struct {
	slice []int
}

func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}
func (s *Stack) Pop() int {
	pop := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return pop
}
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

func (s Stack) String() string {
	return fmt.Sprint(s.slice)
}

// Go language is strongly typed which means you need declare a type implicitly or explicitly
// It's strict when it comes to types and will report errors during compilation.

// Constants can be untyped and you can automatically cast them when assigned.
type customString string
type AnotherString string

const (
	cstring               = "my string"
	cint                  = 1
	cstring2 customString = "hello"
)

type myInt32 int32

func ExampleConversions() {
	// implicit converions
	var b byte
	var ui8 uint8 = 1
	b = ui8
	fmt.Println("byte:", reflect.TypeOf(b).Name(), "|| uint8:", reflect.TypeOf(ui8).Name())
	var r rune
	var i32 int32 = 1
	r = i32
	fmt.Println("rune:", reflect.TypeOf(r).Name(), "|| int32:", reflect.TypeOf(i32).Name())
	var sb []byte
	var sui8 []uint8 = []uint8{1}
	sb = sui8
	fmt.Println("[]byte:", reflect.TypeOf(sb[0]).Name(), "|| []uint8:", reflect.TypeOf(sui8[0]).Name())

	// []int, IntSlice and MySlice share
	// the same underlying type: []int
	type IntSlice []int
	type MySlice []int

	var s = []int{} // s is alias of []int{}
	var is = IntSlice{1}
	var ms = MySlice{1}

	// Must use explicit conversions of IntSlice{} and MySlice{}
	is = IntSlice(ms)
	ms = MySlice(is)

	// Implicit conversions are okay as follows.
	s = is
	is = s
	s = ms
	ms = s
	fmt.Println("MySlice as s:", reflect.TypeOf(ms[0]).Name(), "|| IntSlice as s:", reflect.TypeOf(is[0]).Name())
	// struct aliasing
	var x struct {
		n int `json:"foo"`
	}
	var y struct {
		n int `json:"bar"`
	}
	x = struct {
		n int `json:"foo"`
	}(y)
	y = struct {
		n int `json:"bar"`
	}(x)
}

func ExamplePointerConversion() {
	type MyInt int
	type IntPtr *int
	type MyIntPtr *MyInt

	var pi = new(int)
	var ip IntPtr = pi
	// cannot implicitly convert as var _ *MyInt = pi
	var _ = (*MyInt)(pi) // must be done expliciltly
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

func ExampleStack() {
	s := new(Stack)
	s.Push(100)
	s.Push(22)
	s.Push(37)
	s.Push(54)
	fmt.Println(s)
	fmt.Println(s.Peek())
	fmt.Println("Pop", s.Pop())
	fmt.Println(s)
	fmt.Println("Pop", s.Pop())
	fmt.Println(s.Peek())
	fmt.Println(s)

	// If an underlying type is the same then conversion is valid
	type underlying string
	var us string = "hello"
	under := underlying(us)
	fmt.Println("underlying", reflect.TypeOf(under).Name())

	// however
	// us = underlying("string")
	// will not work because us is already declared a string and not of type underlying
	// another example would be map[string]int =/= map[mystringtype]int where mystringtype of type string
	// you would need to convert mystringtype to a string before using it within the map

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

	var mn customString
	mn = cstring

	var myStringy string
	myStringy = string(cstring2)

	var mfloat float64
	mfloat = cint

	fmt.Printf("asdf %s, %f %s", mn, mfloat, myStringy)
	fmt.Println("reflect", reflect.TypeOf(mn).Name(), reflect.TypeOf(mfloat).Name())
	var another AnotherString
	another = cstring
	fmt.Println("another", reflect.TypeOf(another).Name(), reflect.TypeOf(mfloat).Name())

	//	var n int = 1
	//	var p *int
	//	(*int(n))
	// t := 2 * time.Second is the same as t := time.Duration(2) * time.Second since types cannot be mixed and in this case can only multiple time.Duration

}
