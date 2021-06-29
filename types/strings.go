package types

import (
	"fmt"
	"strings"
)

func ExampleStringConcat() {
	list := []string{"Item1", "Item2", "Item3"}
	var sb strings.Builder
	sb.Grow(len(list))

	for _, v := range list {
		fmt.Fprintf(&sb, "%s, ", v)
	}
	result := sb.String()
	result = result[:sb.Len()-2] // -2 because ", " is part of the formated string that needs to be removed at the end
	fmt.Println(result)
}
