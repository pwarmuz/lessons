package types

import "fmt"

func ExampleArrays() {
	v1 := [5]int{0, 1, 2, 3, 4}
	v2 := v1                                   // Arrays are values. Assigning one array to another copies all the elements.
	copyArraySum := func(i [5]int) (sum int) { // The size of an array is part of its type. [5]int is necessary as that is it's type
		for _, v := range i { // if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
			sum += v
		}
		return
	}
	sum := copyArraySum(v2)
	fmt.Println(sum)
}

func ExampleSlice() {
	v1 := []int{0, 1, 2, 3, 4}
	fmt.Println("original", v1)
	v2 := v1[:4]
	fmt.Println(":4", v2)
	v2 = v1[1:3]
	fmt.Println("1:3", v2)
	v2 = v1[4:]
	fmt.Println("3:", v2)
	sumSlice := func(i []int) (sum int) {
		for _, v := range i { // slice does not get copied
			sum += v
		}
		return
	}
	sum := sumSlice(v1)
	fmt.Println("sum of slice ", sum)

	modifySlice := func(i []int) {
		for k, v := range i { // slice does not get copied, acts like a pointer
			i[k] = v + 1 // the argument will get modified
		}
	}
	modifySlice(v1)
	fmt.Println("modify slice", v1)
}

func ExampleSlicing() {
	five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	five = five[:3]          // five is modified but the original slice still exists
	for _, v := range five { // five is a copy of the most recent assignment so only 0,1,2 index are available, which is what is ranged
		five = five[:2] // re-cutting, will have no affect on the range because it is a copy
		fmt.Printf("v[%s]\n", v)
	}
	/*  This would not work because you are referencing an index that is out of bounds due to the resizing
	for i := range five{
		five = five[:2]  // resizing that will cause index 3 from disappearing and cause a panic in the next line
		fmt.Printf("v[%s]\n", five[i]) // five[3] will be called because i will iterate up to 3 and cause a panic out of range
	}

	*/
	five = five[:5]          // five is re-ranged to it's former limit so all values will be seen
	for _, v := range five { // ranges the most recent value of five
		fmt.Println(v) // all 5 elements will print
	}
}

type Dog struct {
	Name string
	Age  int
}

func ExampleSliceAddress() {
	// this shows how jackie, dogs[0], dogs[1], and dog := range dogs has unique memory addresses
	// dog := range, the dog will always keep the same address because it just copies whatever iterative instance is available from dogs
	jackie := Dog{
		Name: "Jackie",
		Age:  19,
	}

	fmt.Printf("Jackie Addr: %p\n", &jackie) // we are looking for the address

	sammy := Dog{
		Name: "Sammy",
		Age:  10,
	}

	fmt.Printf("Sammy Addr: %p\n", &sammy)

	dogs := []Dog{jackie, sammy}
	fmt.Printf("Jackie Addr: %p\n", &dogs[0])
	fmt.Printf("Sammy Addr: %p\n", &dogs[1])
	fmt.Println("")

	for i, dog := range dogs { // dog will maintain the same memory address and dogs is a copy of the initialization
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n", &dog)
		fmt.Printf("Addr: %p\n", &dogs[i])
		fmt.Println("")
	}
}
func ExampleSliceAddressPtrs() {
	// this does not act the same way as the non pointer version above
	// go passes everything in value but this time the value is a pointer
	// jackie is now a pointer to Dog
	jackie := &Dog{
		Name: "Jackie",
		Age:  19,
	}

	fmt.Printf("Jackie value: %p Addr: %p\n", jackie, &jackie) // in this case we care about the value, which is the pointer

	sammy := &Dog{
		Name: "Sammy",
		Age:  10,
	}

	fmt.Printf("Sammy value: %p, Addr: %p\n", sammy, &sammy)

	dogs := []*Dog{jackie, sammy} // notice that this is a slice of pointers to Dog, Jackie maintains the same pointed address
	fmt.Printf("Jackie value slice: %p, Addr dogs[] %p\n", dogs[0], &dogs[0])
	fmt.Printf("Sammy value slice: %p, Addr dogs[] %p \n", dogs[1], &dogs[1])
	fmt.Println("")

	for i, dog := range dogs {
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Value dog: %p\n", dog)
		fmt.Printf("Addr &dog: %p\n", &dog)
		fmt.Printf("Value dogs: %p\n", dogs)
		fmt.Printf("Addr &dogs: %p\n", &dogs)
		fmt.Printf("Value dogs[]: %p\n", dogs[i])
		fmt.Printf("Addr &dogs[]: %p\n", &dogs[i])
		fmt.Println("")
	}
}

func ExampleMap() {
	timeZone := map[string]int{ // map composite literal
		"UTC": 0,
		"EST": -5,
	}

	offset := func(tz string) (seconds int) {
		seconds, ok := timeZone[tz]
		if !ok {
			fmt.Println("unknown time zone:", tz)
		}
		return
	}
	tz := offset("UTC")
	fmt.Println(tz)
}

type T struct {
	a int
	b float64
	c string
}

// String from Stringer interface
func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

type MyString string

// String note the receiver is directly a string and when used for formating will create a recursive situation
func (s MyString) String() string {
	return fmt.Sprintf("%q\n", string(s)) // Needs to convert s from the receiver as a string or will be recursive
}

func ExamplePrinting() {
	timeZone := map[string]int{ // map composite literal
		"UTC": 0,
		"EST": -5,
	}

	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)      // + annotates the fields of the structure
	fmt.Printf("%#v\n", t)      // # prints full go syntax
	fmt.Println("String() ", t) // # prints based on the String() implementation
	fmt.Printf("%#v\n", timeZone)
	fmt.Printf("%T\n", timeZone) // prints Type

	st := "My string"
	stb := []byte("My bytes")
	si := 1
	sr := rune('A')
	fmt.Printf("%q\n", st)
	fmt.Printf("%#q\n", st)
	fmt.Printf("%q\n", stb)
	fmt.Printf("%#q\n", stb)
	fmt.Printf("%q\n", si)
	fmt.Printf("%#q\n", si)
	fmt.Printf("%q\n", sr)
	fmt.Printf("%#q\n", sr)
}
