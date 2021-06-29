package types

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type ByteSize float64

func TypeInformation() {
	// bitwise XOR (exclusive or) to set each bit to 1
	// because XOR will invert uint(0) of 0 value into a max value filled by 1
	/*
			2^7 | 2^6 | 2^5 | 2^4 | 2^3 | 2^2 | 2^1 | 2^0
		  |  1 |  0  |  1  |  0  |  1  |  0  |  1  |  0  |  Binary representation
			128 +  0  + 32  +  0  +  8  +  0  +  2  +  0  =  170 Sum
	*/
	const (
		MaxUInt8  = ^uint8(0)  // equivalent to 1<<8 - 1
		MaxUInt16 = ^uint16(0) // equivalent to 1<<16 - 1
		MaxUInt32 = ^uint32(0) // equivalent to 1<<32 - 1
		MaxUInt64 = ^uint64(0) // equivalent to 1<<64 - 1

		// convert the unsigned versions into signed through bitwise signed right shift
		MaxInt8 = int8(MaxUInt8 >> 1)
		// take the negative max signed version and subtract 1 (add a negative 1)
		MinInt8  = -MaxInt8 - 1
		MaxInt16 = int16(MaxUInt16 >> 1)
		MinInt16 = -MaxInt16 - 1
		MaxInt32 = int32(MaxUInt32 >> 1)
		MinInt32 = -MaxInt32 - 1
		MaxInt64 = int64(MaxUInt64 >> 1)
		MinInt64 = -MaxInt64 - 1

		_           = iota // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
	)

	fmt.Println("max uInt8", MaxUInt8)
	fmt.Println("max uInt16", MaxUInt16)
	fmt.Println("max uInt32", MaxUInt32)
	fmt.Println("max uInt64", MaxUInt64)

	fmt.Println(MinInt8, " min/ int8 /max ", MaxInt8)
	fmt.Println(MinInt16, " min/ int16 /max ", MaxInt16)
	fmt.Println(MinInt32, " min/ int32 /max ", MaxInt32)
	fmt.Println(MinInt64, " min/ int64 /max ", MaxInt64)

	// Important lesson to note
	// var intValue int = 321 is equivalent to int32(intValue)
	//

	// Bitwise Even/Odd manipulation
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100)
	if number&1 == 1 { // bitwise & operator tests if number is odd
		fmt.Printf("%d is odd\n", number)
	}
	if number&1 == 0 { // bitwise & operator tests if number is even
		fmt.Printf("%d is even\n", number)
	}
	// this works because the number is odd when the least significant bit (&1) is set and it equals 1.
	// when the least significant bit (&1) is set and it equals 0 it means it's even.

}

func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func MinAlt(a ...int) int {
	min := a[0]
	fmt.Println("selected", min)
	for k, i := range a[1:] { // excludes the first entry
		fmt.Println("modified key", k, "v", i)
		if i < min {
			min = i
		}
	}
	return min
}

func ExampleMinInts() {
	series := []int{5, 4, 1, 5, 1, 3, 6, 7, 2, 6, 9}
	min := MinAlt(series...)
	fmt.Println(min)
}
func IntBehavior() {

}

type smallest struct{}

func (s smallest) String() string {
	return fmt.Sprint("struct of 0 size.")
}

func ExampleSmallestType() {

	/*
		   Empty struct struct{} is realized in special way in Go.
			It’s a smallest building block in Go. It’s size is literally 0 bytes.

			If has zero size. you may create a slice of 1000’s empty structures and this slice will be very tiny. Because really Go stores only a number of them in the slice but not them itself. The same story with channels.

			All pointers to it always point to the same special place in memory.

			Very useful In channels when you have notify about some event but you don’t need to pass any information about it, only a fact. Best solution is to pass an empty structure because it will only increment a counter in the channel but not assign memory, copy elements end so on. Sometime people use Boolean values for this purpose, but it’s much worse.

			Zero size container for methods. You may want have a mock for testing interfaces. often you don’t need data on it just methods with predefined input and output.

			Go has no Set object. Bit can be easily realized as a map[keyType]struct{}. This way map keeps only keys and no values.
	*/
	var smallMap map[string]struct{}
	smallMap = make(map[string]struct{})
	smallMap["new"] = struct{}{}
	smallMap["another"] = struct{}{}
	for k := range smallMap {
		fmt.Println("smallmap key", k)
	}
	var smallStruct smallest
	fmt.Println(smallStruct)

	// check out concurrency patters for channels of empty struct
}

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	slice = append(slice, data...)
	*p = slice
	return len(data), nil
}

func ExampleBytes() {
	var b ByteSlice
	n, err := fmt.Fprintf(&b, "This hour has %d days\n", 7)
	if err != nil {

	}
	c := []byte{1, 2, 3}
	b.Write(c)
	fmt.Println(c, n)
}

type Embedded1 struct {
	MyInt   int
	MyFloat float64
}

type Embedded2 struct {
	MySlice []int
	MyMap   map[string]int
}

type Embedder struct {
	*Embedded1
	*Embedded2
}

func (e *Embedded2) Embed() {
	e.MySlice = make([]int, 5)
	e.MyMap = make(map[string]int)
}

func (e *Embedder) Mapped(s string) int {
	e.MyMap = map[string]int{}
	return e.MyMap[s]
}
func (e *Embedder) Mapper(s string) {
	e.MyMap = map[string]int{}
	e.MyMap[s]++
}

func ExampleEmbed() {
	var mine Embedder
	mine.Mapper("hi")
	fmt.Println(mine.Mapped("hi"))
}

func ExampleConst() {
	floater := func(i interface{}) {
		_, ok := i.(float64)
		if ok {
			fmt.Println("Is float64")
		}
		_, ok = i.(float32)
		if ok {
			fmt.Println("Is float32")
		}
	}
	const someFloat64 = 64.0
	floater(someFloat64)
	const someFloat32 = float32(32.0)
	floater(someFloat32)
	const someOtherFloat32 float32 = 32.2
	floater(someFloat32)
	fmt.Printf("%8.4f", someOtherFloat32) // width of 8 units, precision of 4

	const stringFloat64 = "64.064"
	s, err := strconv.ParseFloat(stringFloat64, 64)
	s = s + someFloat64
	if err == nil {
		fmt.Println(s) // 128.064
	}

	checkString := func(i interface{}) {
		_, ok := i.(string)
		if ok {
			fmt.Println("Is string")
			return
		}
		fmt.Println("Not string")
	}
	// default type determined by syntax
	const hello = "hello untyped string constant" //an untyped constant has a default type, an implicit type that it transfers to a value if a type is needed
	const Hello string = "Hello string constant"
	checkString(hello)
	checkString(Hello)

	type fakeString string
	var fs fakeString
	fs = hello
	// fs = Hello // cannot be used as type fakeString because Hello is declared a string string
	fmt.Println(fs)
	checkString(fs)        // not a string
	fs = fakeString(Hello) // explicit conversion
	checkString(fs)        // not a string
}
