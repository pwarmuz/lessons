package structs

import "fmt"

type arithmetic struct {
	a, b int
}

func NewMaths(a, b int) *arithmetic {
	fmt.Println("----- new maths")
	return &arithmetic{a, b}
}
func (ar *arithmetic) MultiplicationPtr() {
	fmt.Println("pre pointer reciever a ", ar.a, "b ", ar.b)
	ar.a = 5
	ar.b = 5
	sum := ar.a * ar.b
	fmt.Println("multiplication ptr", sum, " a ", ar.a, "b ", ar.b)
}
func (ar arithmetic) Multiplication() {
	fmt.Println("pre reciever a ", ar.a, "b ", ar.b)
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
