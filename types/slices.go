package types

import (
	"fmt"
	"unicode/utf8"
)

func ExampleSlices() {
	// index will always start with 0
	// index can be manipulated even with math or bitwise
	numbers := []string{"a", "b", 1 + 4: "c", "d", 9: "e"} // slice literal
	for i, v := range numbers {
		fmt.Println(i, v)
	}
	// var keys []int this is nil
	keys := make([]int, 3) // initializes slice of 3 int zero values
	for i, v := range keys {
		fmt.Println(i, v)
	}
}

func ExampleByteSlice() {
	bs := []byte{71, 111}  // byte slice literal
	fmt.Printf("%s\n", bs) // Output: Go

	sbs := string(bs) // sbs is a copy of bs since the conversion to string creates the copy of []byte
	bs[0] = 55
	fmt.Println(sbs) // because strings are immutable bs and sbs have 2 different memory allocations

	s := "Literal String"
	bs = []byte(s)
	fmt.Printf("%s\n", bs)
	fmt.Printf("%d\n", bs) // decimal value of each byte

	bs = []byte("◺")
	fmt.Println(bs) // Output: [226 151 186]
	s = string(bs)
	fmt.Println("len string", len(s), "len []byte", len(bs)) // Output: 3
	fmt.Println("Rune count", utf8.RuneCountInString(s))     // Output: 1

	for i, b := range "Hi ◺ there" { // remember the ◺ has a length of 3 and will occupy 3 values of the index 3 - 6
		fmt.Printf("i: %d. b: %q\n", i, b) // %q is single character quoted e.g. 'H'. %c doesn't use quotes
	}
}
