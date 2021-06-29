package functions

import (
	"fmt"
	structs "lessons/modules/structs"
	"sort"
)

// Go treats funcions as first class citizens so they can be assigned to variables
// passed as arguments to other fuctions, and returned from other fuctions
func processSlice(i []int, f func(int) int) []int {
	r := make([]int, 0)
	for _, value := range i {
		r = append(r, f(value))
	}
	return r
}

func ExampleProcessSlice() {
	p := []int{2, 3, 5, 8, 13, 26}
	r := processSlice(p, func(n int) int {
		return n + 5
	})
	fmt.Printf("ExampleProcessSlice original %d, processed %d of %T\n", p, r, r)
}

// type alias
type students = structs.Students

func filter(s []students, f func(students) bool) []students {
	m := make([]students, 0)
	for _, value := range s {
		if f(value) {
			m = append(m, value)
		}
	}
	return m
}
func ExampleFilterSlice() {
	s1 := structs.Student1()
	s2 := structs.Student2()
	s := []students{s1, s2}
	f := filter(s, func(sf students) bool {
		if sf.Grade > 3.5 {
			return true
		}
		return false
	})
	fmt.Println(f)
}

func searchInt(sp []int) func(i int) int {
	// sp gets inherited by the closure function which is why it gets used in it's return
	return func(g int) int {
		index := sort.Search(len(sp), func(i int) bool { return sp[i] >= g })
		return index
	}
}
func ExampleSearchInt() {
	p := []int{2, 3, 5, 8, 13, 26}
	si := searchInt(p)
	greater := 6
	v := si(greater)
	fmt.Printf("found %d which is >= %d at index %d\n", p[v], greater, v)
}

func Expand(slice []int, elements ...int) []int {
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]int, n, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}

type password string
type myint int

func ExampleSliceExpand() {
	t := []int{1, 2, 3, 4}
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	fmt.Println("s interface", s)

	pass := password("passy")
	pass = "pissy"
	fmt.Println(pass)

	inty := myint(6)
	inty = 69
	fmt.Println(inty)

	slice := make([]int, 0, 3)
	//additional := []int{6, 9}
	for i := 0; i < 22; i++ {
		slice = Expand(slice, i)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
		fmt.Println("address of 0th element:", &slice[0])
	}
}

func ExampleAnonymous() {
	// anonymous function which accepts an int argument and returns an int value
	fn := func(i int) int { return i * 2 }
	// called below and 6 is used as an argument
	fmt.Println(fn(6))

	// anonymous function where 4 is already assigned to i
	fn2 := func(i int) int { return i * 2 }(4)
	// called below but no argument is used below because it was already declared
	fmt.Println(fn2)

}

// maths function declaration
// function name is maths
// function parameter list is int, int where a and b are formal parameters
// function return types are int
type maths func(a, b int) int

func mathsMultiply() func(a, b int) int {
	// this is a go closure
	return func(a, b int) int {
		// this is the function definition
		// which provides the body of the function
		return a * b
	}
}

func ExampleFuncReturns() {
	multiply := mathsMultiply()
	fmt.Println(multiply(3, 7))

	// Notice the naming schemes of these functions below
	// and the option that anonymous functions allow
	var addition maths = func(a, b int) int {
		return a + b
	}
	fmt.Println("Sum Addition", addition(2, 4))

	// which is essentially this
	subtraction := maths(func(a, b int) int {
		return b - a
	})
	fmt.Println("Sum Subtraction", subtraction(2, 4))

	// You might wonder why even use maths
	division := func(a, b int) int { return b / a }
	fmt.Println("Sum Subtraction", division(2, 4))

	// but you might do something like this
	math := func(m maths, a, b int) int {
		return m(a, b)
	}
	math(subtraction, 2, 4)
	math(addition, 2, 4)
	// notice that division will work
	math(division, 2, 4)
	// because func(int, int) int is equilvalent to maths

}

func callByValue(a int) {
	a *= 5
	fmt.Println("inside call by value", a)
}
func callByReference(a *int) {
	*a *= 5
	fmt.Println("inside call by reference", &a)
}

func callStringByReference(a *string) {
	// this is not a good way to concatinate strings
	// because you are creating a new string which costs run-time
	// but for this purpose of the example...
	*a = *a + " additional"
}

type intMap map[int]int

func callMapByValue(a intMap) {
	// since maps are references you don't need to pass them by reference
	a[3] = 3
}
func callMapByReference(a *intMap) {
	// this is equivalent to callMapByValue since maps are references
	// this is unncessary
	(*a)[4] = 4
}

func ExampleByValueAndReference() {
	var value int = 8
	callByValue(value)
	callByReference(&value)
	var mystring string = "hello"
	callStringByReference(&mystring)
	fmt.Println(mystring)

	myMap := make(intMap, 0)
	myMap[1] = 1
	myMap[2] = 2
	// myMap is being passed by value
	callMapByValue(myMap)
	fmt.Println(myMap[3])
	// & represents reference, &myMap is being passed by reference
	callMapByReference(&myMap)
	fmt.Println(myMap[4])
	// myMap[3] and myMap[4] show that myMap is able to amend
	// and retain information in both situations
	// therefore it is unnecessary to pass maps via reference
}

// namedResultParameter will accept a location name and return the latitude and longitude of the location or an error
// this scheme with naming return parameter helps clarify what the function does
func namedResultParameter(location string) (lat, long int, err error) {
	loc := len(location) // nonsense logic to demonstrate
	lat = 1 * loc        // these below values are returned fulfilling the named results
	long = 2 * loc
	err = nil
	return // this part is a naked return
}
