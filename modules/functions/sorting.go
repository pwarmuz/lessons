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

/*
	for i:=int32(1); i <= n; i++{
		if i %3==0 && i %5 ==0{
			fmt.Println("FizzBuzz")
		}else if i %3 ==0 {
			fmt.Println("Fizz")
		}else if i %5 ==0{
			fmt.Println("Buzz")
		}else{
			fmt.Println(i)
      }
	}

	for i:=0; i<len(fixed)-1; i++{
      reversed[i] = fixed[len(fixed)-i]
	}
*/

        //fmt.Println("row", row)
        lastVal := row[1]
        if row[1]>n{
            lastVal = n
        }
        for i:= row[0]; i<= lastVal; i++ {
            summer[i-1] += row[len(row)-1]
            //fmt.Println("summer", summer)
            if summer[i-1] >= max{
                max = summer[i-1]
            }
        }


		      summer := make([]int32, n+1)

    for _, row := range queries{
        summer[row[0]-1] += row[2]
        summer[row[1]] -= row[2]
    }

    var max int32 = -1
    var total int32 = 0
    for _, v := range summer{
        total += v
        fmt.Println(total)
        if total > max {
            max = total
        }
        if total < 0{
            break
        }
    }
    return int64(max)