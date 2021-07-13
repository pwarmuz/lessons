/*
Lessons will execute any relevant lesson that the programmer wants to see
Since this comment is a public comment it should be written in proper english with capitalization and punctuation.
Use these code review comments for guidance in structure and proper go idioms: https://github.com/golang/go/wiki/CodeReviewComments
*/
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
	"lessons/modules/algorithms"
)

// Intro Go is a general purpose language intended for systems programming.
// It is strongly typed.
// It is statically typed because the types cannot be changed and is most similar to structural typed because it happens at compile time
//
// When writing packages include examples https://blog.golang.org/examples as part of documentation
//
// OO concepts
// Encapsulation possible at the package level
// -- Local retention, protection, and hiding possible with public/private attribute (lower case, capital first letters) and methods
// Composition through embedding
// Polymorphism by satisfying Interfaces
// Inheritance is NOT possible
// Messaging with Channels
// Late-binding possible via higher-order-function (function that takes a function as an argument or returns function) and interfaces
func main() {
	//fmt.Println("Hello")
	//fmt.Println(sample.SillyName())
	//functions.ExampleProcessSlice()
	//functions.ExampleFilterSlice()
	fmt.Println("...")
	//functions.ExampleMiddleware()
	//functions.ExampleSearchInt()
	//functions.ExampleSliceExpand()
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
	//conversions.ExampleNumericConversions()
	//functions.ExampleByValueAndReference()
	//concurrency.ExampleUnbufferedChan()
	//concurrency.ExampleBufferedChan()
	//concurrency.ExUnbufferedChan()
	//concurrency.ExTimeOuts()
	//concurrency.ExampleSimpleTimer()
	//concurrency.ExampleSimpleTicker()
	//concurrency.ExWorkerPool()
	//types.TypeInformation()
	//algorithms.ExampleAlgoPatterns()
	//interfaces.ExampleInterface()
	//types.ExampleSlice()
	//types.ExamplePrinting()
	//types.ExampleMinInts()
	//types.ExampleBytes()
	//interfaces.ExampleRot13NoStruct()
	//types.ExampleEmbed()
	//interfaces.ExampleAssertion()
	//interfaces.ExampleErrorPanicRecovery()
	//structs.ExamplePointerVsValue()
	//structs.ExampleImplementation()
	//types.ExampleSliceAddressPtrs()

	//algorithms.ExampleRecursiveFunctions()
	//algorithms.ExamplePermutations()
	algorithms.ExampleCommandPattern()
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
