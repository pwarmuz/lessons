package functions

import "fmt"

func addVariadic(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
func ExampleVariadic() {
	one := 1
	two := 2
	three := 3
	fmt.Println("adding...", one, two, three, "and 4; sum: ", addVariadic(one, two, three, 4))

	addThis := []int{1, 2, 3, 4, 5}
	fmt.Println("adding...", addThis, "; sum: ", addVariadic(addThis...))

	// cannot combine the "addThis..." with previous entries such as 'addVariadic(one, two, three, 4, addThis...))'. This will not work
}
