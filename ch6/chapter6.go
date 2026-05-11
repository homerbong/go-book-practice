package main

import (
	"fmt"
	"go-book-practice/display"
)

type Foo struct {
	foo string
	bar int
}

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func makePointer[T any](t T) *T {
	return &t
}

func pointerSyntaxExamples() {
	x := "Hello"
	pointerToX := &x

	fmt.Println("x:", x)
	fmt.Println("pointerToX:", pointerToX)

	y := 10
	pointerToY := &y
	fmt.Println("y:", y)
	fmt.Println("pointerToY:", pointerToY)
	fmt.Println("*pointerToY:", *pointerToY)

	z := 5 + *pointerToY
	fmt.Println("z:", z)
}

func pointerTypeExamples() {
	var x *int
	fmt.Println("x == nil:", x == nil)
	// Can't dereference before assigning
	// fmt.Println("*x:", *x) // panic: runtime error: invalid memory address or nil pointer dereference

	y := 10
	var pointerToY *int
	pointerToY = &y
	fmt.Println("pointerToY:", pointerToY)
	fmt.Println()

	var z = new(int) // Instantiates a pointer to a zero value instance of the type.
	fmt.Println("z != nil:", z != nil)
	fmt.Println("*z:", *z)
	fmt.Println()

	str1 := struct {
		name string
		age  int
	}{}
	fmt.Println("str:", str1)
	pointerToStruct := &struct {
		name string
		age  int
	}{}
	fmt.Println("pointerToStruct:", pointerToStruct)
	fmt.Println("*pointerToStruct:", *pointerToStruct)
	fmt.Println()

	foo := &Foo{}
	fmt.Println("foo:", foo)
	fmt.Println("*foo:", *foo)
	fmt.Println()

	var primitive string
	pointerToPrimitive := &primitive
	fmt.Println("primitive:", primitive)
	fmt.Println("pointerToPrimitive:", pointerToPrimitive)
	fmt.Println()

	// p := person {
	// 	FirstName:  "Alex",
	// 	MiddleName: "Perry", // build error: cannot use "Perry" (untyped string constant) as *string value in struct literal
	// 	LastName:   "Peterson",
	// }

	// p := person{
	// 	FirstName:  "Alex",
	// 	MiddleName: &"Perry", // build error: invalid operation: cannot take address of "Perry" (untyped string constant)
	// 	LastName:   "Peterson",
	// }

	// Alternative 1:
	p := person{
		FirstName:  "Alex",
		MiddleName: pointerToPrimitive,
		LastName:   "Peterson",
	}
	fmt.Println("p:", p)

	// Alternative 2:
	p = person{
		FirstName:  "Alex",
		MiddleName: makePointer("Perry"),
		LastName:   "Peterson",
	}
	fmt.Println("p:", p)
}

func failedUpdate(g *int) {
	fmt.Println("g inside before reassinging:", g)
	// The function receives a copy of the pointer so if it's a nil it won't be updated.
	x := 10
	g = &x
	fmt.Println("g inside failedUpdate:", g)
}

func update(px *int) {
	*px = 55
}

func reassingingPointersExamples() {
	var x int = 10
	var y int = x
	y = 20
	fmt.Println("x y:", x, y)
	fmt.Println()

	var f *int // nil pointers can't be updated
	fmt.Println("f before failedUpdate:", f)
	failedUpdate(f)
	fmt.Println("f after failedUpdate:", f)
	fmt.Println()

	val := 5
	f = &val
	fmt.Println("f before failedUpdate:", f)
	fmt.Println("*f before failedUpdate:", *f)
	fmt.Println()
	failedUpdate(f)
	fmt.Println()                             // This changes the value of the pointer of the internal variable.
	fmt.Println("f after failedUpdate:", f)   // It does not change the pointer itself.
	fmt.Println("*f after failedUpdate:", *f) // It does not change the value to which the pointer points to.
	fmt.Println()
	update(f)                           // This one instead changes the value to which the pointer points to
	fmt.Println("f after Update:", f)   // The pointer itself does not changes since it's copied.
	fmt.Println("*f after Update:", *f) // The value does change here
}

func pointersExamples() {
	pointerSyntaxExamples()
	fmt.Println()

	pointerTypeExamples()
	fmt.Println()

	reassingingPointersExamples()
	fmt.Println()
}

func main() {
	display.SectionTitle("Pointers")
	pointersExamples()
}
