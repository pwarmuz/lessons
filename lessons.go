package main

// To upgrade project to v2/
// Create v2/ mkdir v2  THEN $ cp go.mod v2/go.mod  Copying contents into new version if needed
// Copy the go mode file to v2/ cp go.mod /v2/go.mod  THEN  $ go mod edit -module {path/v2} v2/go.mod
// go mod tidy
// go mod why -m golang.org/x/text
// standard Caddy modules
//_ "www-phaeton/v2/caddyconfig/caddyfile"
//_ "www-phaeton/v2/modules/caddyhttp/standard"
//_ "www-phaeton/v2/modules/filestorage"
//_ "www-phaeton/v1/modules/logging"
// git tag -l v1.0.0-b* or git tag -l *beta*
import (
	"fmt"
	"lessons/modules/concurrency"
)

func main() {
	//fmt.Println("Hello")
	//fmt.Println(sample.SillyName())
	//functions.ExampleProcessSlice()
	//functions.ExampleFilterSlice()
	fmt.Println("...")
	//functions.ExampleMiddleware()
	//functions.ExampleSearchInt()
	//functions.ExampleSliceExpand()
	//functions.ExampleMiddleBool()
	//structs.MethodMath()
	//functions.ExampleVariadic()
	//concurrency.ExampleCurrency()
	//algorithms.ExampleQueue()
	//algorithms.ExampleAltQueue()
	//algorithms.ExampleStack()
	//algorithms.ExampleConversions()
	//functions.SortingSlice()
	//conversions.ExampleConversionCosts()
	//functions.ExampleByValueAndReference()
	concurrency.ExampleUnbufferedChan()
	concurrency.ExampleBufferedChan()
	concurrency.ExampleBufferedChanRoutine()

}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
