package functions

import (
	"fmt"
	"time"
)

// closures can be used to:
// -isolate data
// -wrapping functions to create middleware
// -access data without needing to create global variables via middleware
//
func timingPtr(f func(*[]int) *[]int) func(*[]int) *[]int {
	return func(si *[]int) *[]int {
		start := time.Now()
		sliced := f(si)
		end := time.Now()
		fmt.Println("ptr process time took", end.Sub(start))
		return sliced
	}
}

// Slices are sliceheaders which have a pointer to the location
// timingPtr and timingCopy are the same because the pointer is being accessed
// and slices are references to that pointer
func timingCopy(f func([]int) []int) func([]int) []int {
	return func(si []int) []int {
		start := time.Now()
		sliced := f(si)
		end := time.Now()
		fmt.Println("copy process time took", end.Sub(start))
		return sliced
	}
}

func timingBool(f func([]int) []int) func([]int) bool {
	return func(si []int) bool {
		start := time.Now()
		sliced := f(si)
		end := time.Now()
		fmt.Println("copy process time took", end.Sub(start), " for ", sliced)
		return 0 < len(sliced)
	}
}

// processSlicerPtr will re-iterate the point that slices are references to pointers
func processSlicerPtr(iSlices *[]int) *[]int {
	for i, value := range *iSlices {
		(*iSlices)[i] = value * 3
	}
	return iSlices
}

func processSlicerCopy(iSlices []int) []int {
	for i, value := range iSlices {
		iSlices[i] = value * 3
	}
	return iSlices
}

func ExampleMiddleware() {
	p := []int{2, 3, 5, 8, 13, 26}
	// d := func(si []int) func([]int) {
	// 	for i, value := range si {
	// 		si[i] = value * 3
	// 	}
	// 	fmt.Println("")
	// 	return func(si []int) {
	// 		fmt.Println("inside D", si)
	// 	}
	// }(p)

	d1 := timingCopy(processSlicerCopy)
	slicedCopy := d1(p)

	d2 := timingPtr(processSlicerPtr)
	slicedPtr := d2(&p)

	fmt.Println("timed with copy", slicedCopy)
	fmt.Println("timed with pointers", *slicedPtr)
}

func ExampleMiddleBool() {
	p := []int{2, 3, 5, 8, 13, 26}
	sl := timingBool(processSlicerCopy)
	if sl(p) {
		fmt.Println("Timing finished")
	}
}
