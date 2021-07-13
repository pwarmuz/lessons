package structs

import (
	"fmt"
	"lessons/modules/interfaces"
	"unsafe"
)

// When you pass a value to a function or method, a copy is made and passed.
// If you pass a string to a function, only the structure describing the string will be copied and passed, since strings are immutable.
// Same is true if you pass a slice (slices are also descriptors). Passing a slice will make a copy of the slice descriptor but it will refer to the same underlying array.
type arithmetic struct {
	a, b int
}

func NewMaths(a, b int) *arithmetic {
	fmt.Println("----- new maths")
	return &arithmetic{a, b}
}
func (ar *arithmetic) MultiplicationPtr() {
	fmt.Println("pre pointer receiver a ", ar.a, "b ", ar.b)
	ar.a = 5
	ar.b = 5
	sum := ar.a * ar.b
	fmt.Println("multiplication ptr", sum, " a ", ar.a, "b ", ar.b)
}
func (ar arithmetic) Multiplication() {
	fmt.Println("pre receiver a ", ar.a, "b ", ar.b)
	ar.a = 5
	ar.b = 5
	sum := ar.a * ar.b
	fmt.Println("multiplication ", sum, " a ", ar.a, "b ", ar.b)
}
func (ar arithmetic) Addition() {
	fmt.Println("multiplication pointers")
	ar.MultiplicationPtr()
	fmt.Println("post adjustment a ", ar.a, "b ", ar.b)
	sum := ar.a + ar.b
	fmt.Println("addition ", sum, " a ", ar.a, "b ", ar.b)
	fmt.Println("------------------------------------")
}
func (ar arithmetic) Addition2() {
	fmt.Println("no pointers")
	ar.Multiplication()
	fmt.Println("post multiplication a ", ar.a, "b ", ar.b)
	sum := ar.a + ar.b
	fmt.Println("addition ", sum, " a ", ar.a, "b ", ar.b)
	fmt.Println("------------------------------------")
}
func (ar *arithmetic) Amend(math *arithmetic) {
	fmt.Println("amending...", math.a, math.b)
	*ar = *math
}
func MethodMath() {
	// sync
	ar := arithmetic{8, 9}
	ar.Addition()
	ar.Addition2()
	ar.Addition()
	//ar2 := NewMaths(3, 3)
	ar2 := arithmetic{3, 3}
	ar.Amend(&ar2)
	ar.Addition()
	ar.Addition2()
	ar.Addition()

}

type details struct {
	name    string
	age     int
	balance float64
}

type embeddedDetails struct {
	name string
	age  int
	additional
	a1 additional
	details
	exp interfaces.Exportable
}

// this whole struct is not exportable
type additional struct {
	alias  string
	height int
}

func ExampleImplementation() {
	var zeroValue details
	fmt.Println("Zero values:", zeroValue)
	zeroValue.name = "Sam"
	zeroValue.age = 22
	zeroValue.balance = 123.55
	fmt.Println("Assigned values:", zeroValue)

	compositeLiteral := details{"Steve", 23, 123.43} // values are defined
	fmt.Println("composite without keys:", compositeLiteral)

	compositeLiteral2 := details{name: "Nacho"} // zero values for remaining keys, name needs to be explicit
	fmt.Println("composite with keys and zero value:", compositeLiteral2)

	// a composite literal must either have no keynames but all fields filled out
	compositeLiteralEmbedded := embeddedDetails{
		"Bonus",
		33,
		additional{
			"TopGun",
			55,
		},
		additional{
			"TopGun",
			55,
		},
		details{ // it is ideal to label keys explicitly since it's hard to figure out what is the age or balance
			"asdf",
			33,   // this is age but not easy to see at quick glance
			34.5, // this is the float64 but more attention is required because the key is not present
		},
		interfaces.Exportable{ // outer level has no key names, this inner level does, each level needs to have the appropriate consistent pattern
			PubInt:    55,
			PubString: "Exported",
		},
	}
	fmt.Println("embedded composite external package, outer level no keys:", compositeLiteralEmbedded.exp.PubInt)

	// or a composite literal can have key names with partial fields missing given zero values
	compositeLiteralEmbeddedKeys := embeddedDetails{ // outer level is filled with keys
		name: "Bonus",
		// age:  0, // will be missing and be given a zero value
		additional: additional{ // inner levels all filled with keys
			alias:  "TopGun",
			height: 55,
		},
		a1: additional{
			alias:  "TopGun",
			height: 55,
		},
		details: details{
			name:    "asdf",
			age:     33,
			balance: 44.4,
		},
		exp: interfaces.Exportable{
			PubInt:    55,
			PubString: "Exported",
		},
	}
	fmt.Println("embedded composite external package, outer level with keys:", compositeLiteralEmbeddedKeys.exp.PubInt)
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) String() string {
	// Note Pointer receivers are not concurrent safe because they are mutable
	return fmt.Sprintf("%s is a %s dog breed, from with a pointer receiver.", d.Name, d.Breed)
}

func (d *Dog) Rename(name string) {
	d.Name = name
}

func (d *Dog) Walk(distance string) string {
	return fmt.Sprintf("%s walked %s", d.Name, distance)
}

type Walker interface {
	Walk(string) string
}

type temp struct{}

func (t temp) Walk(d string) string {
	return fmt.Sprint("not walking ", d)
}

func NewWalker(d *Dog) Walker {
	// interfaces are built of concrete value(data) and a dynamic type, the d of *Dog becomes the data of the interface and can be nil without causing the full walker interface to be nil
	// walker's type needs to be nil in order to trigger the "walked == nil" logic in the Walking function
	// in order to do this Walker is returned nil if dog is nil
	if d == nil {
		return nil // ensures the Walker interface is nil
	}
	return d
}
func Walking(walked Walker) string {
	if walked == nil {
		return fmt.Sprintf("Walker is nil value:%v Type:%T", walked, walked)
	}
	return fmt.Sprintf("%s along with the Walking function. value:%v Type:%T", walked.Walk("200"), walked, walked)
}

type Cat struct {
	Name  string
	Breed string
}

func (c Cat) String() string {
	// Note value receivers are concurrent safe because they are immutable due to passing the value
	return fmt.Sprintf("%s is a %s cat breed, from with a value receiver.", c.Name, c.Breed)
}
func (c *Cat) Rename(name string) {
	c.Name = name
}

func ExamplePointerVsValue() {
	// https://npf.io/2014/05/intro-to-go-interfaces/ has some good information on this, most of it is covered here
	dog := Dog{"Beanie", "Water"}
	dog.Rename("Bernie") // since the composite literal is a value and the method has a pointer receiver, Go will treat dog.Rename as (&dog).Rename so it automatically works
	fmt.Println(dog)     // will not print the custom string because the String method has pointer receiver and .String() is not explicitly called
	// As written in Effective Go, The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers. This is because pointer methods can modify the receiver; invoking them on a copy of the value would cause those modifications to be discarded.
	// This this important in relation to fulfilling interfaces, as seen when the stringer interface was not fulfilled because the String method for Dog type used a pointer receiver
	fmt.Println(dog.Walk("500"))
	fmt.Println(Walking(&dog))
	var notADog *Dog // this is a nil Dog struct
	//fmt.Println(Walking(notADog)) // invalid memory run-time panic
	w := NewWalker(notADog)
	fmt.Println(Walking(w)) // Walker is nil
	dogPtr := &dog
	fmt.Println(dogPtr, "pointer as dogPtr")            // works because it's a pointer
	fmt.Println(&dog, "pointer as &dog")                // works because it's a pointer
	fmt.Println(dog.String(), "explicit method called") // works because as mentioned earlier Go will treat this as (&dog).String()
	cat := Cat{"Sanders", "Siamese"}
	cat.Rename("Bonkers")
	fmt.Println(cat, "5")
	fmt.Println(&cat, "5")
	cat2 := &Cat{"Bonkers", "Siamese"}
	cat2.Rename("Bongo")
	fmt.Println(cat2, "5")

	var nilTemp temp
	fmt.Println(nilTemp.Walk("500")) // will print regardless of nil struct, temp's implementation of Walk does not use the struct so nothing is dereferenced so nothing to panic

	// conversion
	var fmtStringer fmt.Stringer
	someDog := &Dog{"Assigned to var fmtStringer", "Water"}
	fmtStringer = someDog       // someDog uses a pointer receiver
	fmt.Println(fmtStringer)    // satisfies fmt.Stringer with a pointer receiver
	fmtStringer = Cat(*someDog) // pointer to someDog can be converted to Cat
	// Cat has a value receiver and can be invoked on a pointers and value receivers. It doesn't matter what Dog actually has because the important thing is Cat, which is converting Dog has a value receiver.
	//fmtStringer.Rename("renamed")  // Cannot use any other methods besides String because fmtStringer is of type Stringer interface not Dog or Cat
	fmt.Println(fmtStringer)

	/*  You cannot convert Dog from Cat
	someCat := &Cat{"Assigned to var fmtStringer", "Water"}
	fmtStringer = someCat    // someCat uses a value receiver
	fmt.Println(fmtStringer) // satisfies the default value receiver
	fmtStringer = Dog(*someCat) // cannot convert someCat into Dog because dog does not have a value receiver.
	***** The important thing is Dog has a pointer receiver and cannot convert Cat regardless of what receiver Cat has.
	***** Both Dog and Cat would need Value receivers to alternate and convert eachother. Only the struct with the value receiver can convert the other struct.
	fmt.Println(fmtStringer)
	*/

	anonFunc := func(i fmt.Stringer) {
		fmt.Println("Anon func", i)
	}
	anonFuncPtr := func(i *fmt.Stringer) { // this *fmt.Stringer is a concrete type. just like []interface{} is a concrete type.
		fmt.Println("Anon func Ptr", *i) // pointer to i is needed to see the string values
	}
	someCat := Cat{"Assigned to var fmtStringer", "Water"}
	anonFunc(someDog)
	anonFunc(someCat)
	anonFunc(fmtStringer) // since fmtString was last converted to a Cat it will be a cat breed
	// anonFunc(Dog(someCat)) // this is analogous to what was mentioned with *****. Cannot use Dog as fmt.Stringer value in argument because missing method String with a value receiver.
	var fmtStringerDogPtr fmt.Stringer = someDog  // someDog needs to be a pointer to fulfil the pointer receiver method
	var fmtStringerCat fmt.Stringer = someCat     //someCat can be a value or pointer and still fulfill the value receiver method
	var fmtStringerCatPtr fmt.Stringer = &someCat //someCat can be a value or pointer and still fulfill the value receiver method
	anonFuncPtr(&fmtStringerDogPtr)               // cannot use someDog unless it is explicitly a fmt.Stringer type, *Dog is assigned to fmtStringerDog as a value and still needs to be dereferenced
	anonFuncPtr(&fmtStringerCat)                  // cannot use someCat unless it is explicitly a fmt.Stringer type, Cat is assigned to fmtStringerCat as a value and still needs to be dereferenced
	anonFuncPtr(&fmtStringerCatPtr)               // cannot use someCat unless it is explicitly a fmt.Stringer type, Cat is assigned to fmtStringerCat as a value and still needs to be dereferenced
}

type admin struct {
	ID int
	Me user
}
type user struct {
	name  string
	likes int
}

func (u *user) notify() {
	fmt.Printf("%s has %d likes\n", u.name, u.likes)
}

func (u *user) addLike() {
	u.likes++
}

func ExamplePointerAndValue() {
	users := []user{
		{name: "bob"},
		{name: "clarissa"},
	}

	// ranging and using the index as a pointer to method notify
	for i := range users {
		users[i].addLike()
	}

	// ranging and using the user value as a value to method notify
	for _, user := range users {
		user.notify()
	}

	admin := admin{
		12345,
		user{
			"john",
			23,
		},
	}
	admin.Me.addLike()
	admin.Me.notify()
}

// emptyStruct has various benefits
// an empty Struct is of size zero and all empty structs point to the same location
// methods can be defined on it
// implement an interface to it since you can use methods
// as a singleton and achieved by:
// 	- store all data in global variables
//    - there is only 1 instance of the type since all empty structs are interchangeable
type emptyStruct struct{}

func (e emptyStruct) Strings() string {
	return "I will always return this"
}

func ExampleSizeOfStruct() {
	// The value's width is always a multiple of it's alignment
	// How it is ordered within the structure IS important in how the memory becomes allocated
	type S1 struct {
		a int8  // 1 (it's bytes)
		b int16 // 2
		c int32 // 4
		d int64 // 8

	}
	/* where xx is empty but needed to fulfil a multiple of it's alignment, it is not a sum of all sizes
	where 08 represents int8 / 16 = int16 / 32 = int32 / 64 = int64
	[08][16][16][32][32][32][32][xx] 8
	[64][64][64][64][64][64][64][64] 16
	*/
	var s1 S1
	fmt.Println(unsafe.Sizeof(s1)) // prints 16

	type S2 struct {
		a int64 // 8 (it's bytes)
		b int64 // 8
		c int64 // 8
		d int64 // 8

	}
	/* where xx is empty but needed to fulfil a multiple of it's alignment, it is not a sum of all sizes
	[64][64][64][64][64][64][64][64] 8
	[64][64][64][64][64][64][64][64] 16
	[64][64][64][64][64][64][64][64] 24
	[64][64][64][64][64][64][64][64] 32
	*/
	var s2 S2
	fmt.Println(unsafe.Sizeof(s2)) // prints 32

	type S3 struct {
		a int32 // 4 (it's bytes)
		b int32 // 4
		c int32 // 4
		d int64 // 8

	}
	/* where xx is empty but needed to fulfil a multiple of it's alignment, it is not a sum of all sizes
	[32][32][32][32][32][32][32][32] 8
	[32][32][32][32][xx][xx][xx][xx] 16
	[64][64][64][64][64][64][64][64] 24
	*/
	var s3 S3
	fmt.Println(unsafe.Sizeof(s3)) // prints 24

	type S4 struct {
		a int32 // 4 (it's bytes)
		b int32 // 4
		c int32 // 4
		d int8  // 1
		e int8  // 1
		f int64 // 8
	}
	/* where xx is empty but needed to fulfil a multiple of it's alignment, it is not a sum of all sizes
	[32][32][32][32][32][32][32][32] 8
	[32][32][32][32][08][08][xx][xx] 16
	[64][64][64][64][64][64][64][64] 24
	*/
	var s4 S4
	fmt.Println(unsafe.Sizeof(s4)) // prints 24

	type S5 struct {
		a int32 // 4 (it's bytes)
		b int32 // 4
		c int32 // 4
		d int64 // 8
		f int8  // 1
		g int8  // 1
		h int16 // 2

	}
	/* where xx is empty but needed to fulfil a multiple of it's alignment, it is not a sum of all sizes
	[32][32][32][32][32][32][32][32] 8
	[32][32][32][32][xx][xx][xx][xx] 16
	[64][64][64][64][64][64][64][64] 24
	[08][08][16][16][xx][xx][xx][xx] 32
	*/
	var s5 S5
	fmt.Println(unsafe.Sizeof(s5)) // prints 32 Notice how the memory allocated between S5 and S6!

	type S6 struct {
		a int32 // 4 (it's bytes)
		b int32 // 4
		c int32 // 4
		d int8  // 1
		e int8  // 1
		f int16 // 2
		g int64 // 8

	}
	/* where xx is empty but needed to fulfil a multiple of it's alignment, it is not a sum of all sizes
	[32][32][32][32][32][32][32][32] 8
	[32][32][32][32][08][08][16][16] 16
	[64][64][64][64][64][64][64][64] 24
	*/
	var s6 S6
	fmt.Println(unsafe.Sizeof(s6)) // prints 24 Notice how the memory allocated between S5 and S6!
}
