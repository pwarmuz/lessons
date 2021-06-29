package algorithms

import "fmt"

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
