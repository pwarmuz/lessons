package algorithms

import (
	"fmt"
	"sort"
)

func PatternFilter(names []string, verify string) bool {
	for _, v := range names {
		if v == verify {
			return true
		}
	}
	return false
}

type Incrmented struct {
	data []int
}

func (d *Incrmented) PatternCompareNIncrement(value int) {
	for v := range d.data {
		if v == value {
			fmt.Println("ignoring", v)
			return
		}
	}
	fmt.Println("adding", value)
	d.data = append(d.data, value)
}

func ExampleAlgoPatterns() {
	var lists Incrmented
	lists.PatternCompareNIncrement(0)
	lists.PatternCompareNIncrement(1)
	lists.PatternCompareNIncrement(1)
	fmt.Println(lists.data)
}

func comparePattern(a, b int) bool {
	if a == b {
		return true
	}
	return false
}
func compareDepthPattern(a, b []byte) int {
	// this handles 2 situations of depth for this comparison
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}
func genericTypeSwitch(t interface{}) {
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}

func compareIncrement(a, b int) int {
	if a == b {
		a++ // if you want to increment while a == b
		return a
	}
	// a++ if you want to increment while a != b
	return a
}
func something() error {
	return nil
}

func multipleShortHandPattern() {
	// this pattern is effective if you have multiple functions using the same variable
	// and don't want to worry about managing the shorthand := form or issues with scope
	var err error
	err = something() // typically a shorthand would be declared here at the start
	// but if the order changed or functions were declared earlier than the shorthand would be edited
	// by declaring var err error at the start you don't need to manage the shorthand form and the scope is
	if err != nil {
		panic(err)
	}
	err = something()
	if err != nil {
		panic(err)
	}
	err = something()
	if err != nil {
		panic(err)
	}
}

func sliceCopying() {
	var nilSlice []int // this is nil
	if nilSlice == nil {
		fmt.Printf("%#v", nilSlice)
	}

	arr := []int{1, 2, 3}
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	fmt.Println("tmp", tmp, "arr", arr)

	emptySlice := []int{}                   // empty slice literal, it's non nil but zero length
	emptySlice = append(emptySlice, arr...) // make sure to use the variadic form
	fmt.Println("emptySlice", emptySlice, "arr", arr)

	arrByValue := arr
	fmt.Println("arrByValue", arrByValue)
	arrByPointer := &arr
	fmt.Println("arrByPointer", *arrByPointer)
	arr[0] = 8
	// remember, slices and maps use pointers to find a node
	// so the original slice was amended but it will reflect the change on both copies because they are referencing the same data
	fmt.Println("arrByValue", arrByValue)
	fmt.Println("arrByPointer", *arrByPointer)
	// however the tmp copy will not have changed
	fmt.Println("tmp", tmp, "arr", arr)

	// copy overwrites a current slice
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := copy(s, s[6:])    // 6,7,8,9 is the slice of data being copied
	fmt.Println(s, "n", n) // 6,7,8,9,4,5,6,7,8,9 ; index 0-4 are copied over
}

func mapsAndKeys() {
	// Option 1) define map & instantiate it
	// var m1 map[int]string
	// m1 = make(map[int]string)

	// Option 2) implicit map instantiation
	// m2 := make(map[int]string)

	// Option 3
	mi := map[int]string{ // map literal
		1: "one",
		2: "two",
		3: "three",
		9: "nine",
	}
	//	need to a stable iteration order, so we need another data structure
	keys := make([]int, len(mi)) // need to assign a length because we are re-assigning a value to the index
	var k int
	for i := range mi {
		keys[k] = i
		fmt.Println("slice key", k, "value", i)
		k++
	}

	sort.Ints(keys) // need to sort because maps are unordered and will

	for i, v := range keys {
		fmt.Println(i, mi[v])
	}
}

// permutateRune the values at index i to len(a)-1.
func permutateRune(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	permutateRune(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permutateRune(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func ExamplePermutations() {
	// Perm calls f with each permutation of a.
	Perm := func(a []rune, f func([]rune)) {
		permutateRune(a, f, 0)
	}

	Perm([]rune("abc"), func(a []rune) {
		fmt.Println(string(a))
	})
}

func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func reverseSlice(si []int) {
	for a, b := 0, len(si)-1; a < b; a, b = a+1, b-1 {
		si[a], si[b] = si[b], si[a]
	}
}

// directRecursion a function that calls itself directly without another function to reach a base condition which is the final executable statement.
// this is also a finite recursion: it calls itself, it stops at a base condition
// a regular recursion does not finish the calculation immediately
func directRecursion(num int) int {
	defer fmt.Println("directRecursion deferred", num)
	if num == 0 || num == 1 {
		return 1
	}
	if num < 0 {
		return -1
	}
	return num * directRecursion(num-1)
}

// infiniteRecursion a function that calls itself directly with not base condition.
func infiniteRecursion() {
	fmt.Println("THIS WILL NEVER END")
	infiniteRecursion()
}

// tailRecursion happens when a function calls itself and calculates a value and sends it down the stack.
func tailRecursion(number, product int) int {
	defer fmt.Println("tailRecursion deferred", number)
	product = product + number
	if number == 0 {
		return product
	}
	return tailRecursion(number-1, product)
}

func tail(num int) {
	defer fmt.Println("tailRecursion deferred without return calls", num)
	if num == 0 {
		fmt.Println("0")
		return
	}
	fmt.Println(num)
	tail(num - 1)
}

func ExampleRecursiveFunctions() {
	/*
			As mentioned in https://www.ardanlabs.com/blog/2013/09/recursion-and-tail-calls-in-go_26.html
		" Nothing has been optimized for the tail call we implemented. We still have all the same stack manipulation and recursive calls being made. So I guess it is true that Go currently does not optimize for recursion. This does not mean we shouldn’t use recursion, just be aware of all the things we learned.

		If you have a problem that could best be solved by recursion but are afraid of blowing out memory, you can always use a channel. Mind you this will be significantly slower but it will work. "

		I think the idea is to just implement recursion as simply and easily to understand as possible since performance difference will be neglegible.
	*/
	var answer int
	answer = directRecursion(0)
	fmt.Println(answer)
	fmt.Println("***ended first recursion")

	answer = directRecursion(5)
	fmt.Println(answer)
	fmt.Println("***ended second recursion")

	answer = tailRecursion(5, 0)
	fmt.Println(answer)
	fmt.Println("***ended third recursion")

	tail(6)
	fmt.Println("***ended fourth recursion")
}
