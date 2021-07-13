package conversions

import (
	"fmt"
	"time"
)

const (
	numeratorf   float64 = (1.5 * 22) + 12
	denominatorf float64 = .25 + 10
	numeratord   int     = (3 * 15) + 12
	denominatord int     = 2*23 + 10
)

func iToF64(i int) float64 {
	return (float64(i) * numeratorf) / denominatorf
}

func fToI(f float64) int {
	return (int(f) * numeratord) / denominatord
}

func ExampleNumericConversions() {
	start := time.Now()
	max := 1500
	for i := 1; i < max; i++ {
		iToF64(i)
	}
	fmt.Println("Final conversion:", iToF64(max), "iterations:", max)
	end := time.Now()
	fmt.Println("Numeric Conversion took", end.Sub(start))
}
