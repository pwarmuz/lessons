package functions

import (
	"fmt"
	"sort"
)

func SortingSlice() {
	girls := []string{"melissa", "elena", "emmanuelle", "clarissa", "ann"}
	sort.Sort(byPriority(girls))
	fmt.Println(girls)
	fmt.Printf("Clarissa should be first, because %s explains it all\n", girls[0])
}

type byPriority []string

func (bp byPriority) Len() int {
	return len(bp)
}
func (bp byPriority) Swap(i, j int) {
	bp[i], bp[j] = bp[j], bp[i]
}
func (bp byPriority) Less(i, j int) bool {
	// clarissa always takes priority
	if bp[i] == "clarissa" {
		return true
	}
	// rest is sorted by length
	// since we want the length that is greatest to be earliest
	// we can compare j to i and order it that way
	// if j is less than i is is therefore "Less" (as named in the method)
	return len(bp[j]) < len(bp[i])
	// if we re-ordered the comparison it would still need to be
	// i > j but would but less pleasant to look at because
	// it would be phrased as i is greater than j which makes j "Less" (as named in the method)
	//return len(bp[i]) > len(bp[j])
	// which makes it cleaner to re-order as above
}
