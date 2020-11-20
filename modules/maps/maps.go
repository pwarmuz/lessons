package maps

import (
	"fmt"
)

// maps are references
// Maps are unordered (aka unsorted)
func ExampleMaps() {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// maps are pointers so
	// m2 points to m1
	var m2 = m1

	// making a change on m2...
	m2["nine"] = 9
	// will mimic that change to m1
	fmt.Println("m1", m1)
	fmt.Println("m2", m2)
}
