package conversions

import (
	"fmt"
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
