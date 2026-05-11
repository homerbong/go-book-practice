package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-book-practice/display"
	"time"
)

type Foo struct {
	foo string
	bar int
}

type FooJson struct {
	Foo string `json:"foo"` // Capital property name on the struct and no spaces between
	Bar int    `json:"bar"`
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

/*
DO NOT USE. WRONG PATTERN!
*/
func makeFooWrong(f *Foo) error {
	if f == nil { // This check is needed to avoid panics.
		return errors.New("Can't assign to a nil pointer.")
	}
	f.foo = "Hello"
	f.bar = 1
	return nil
}

func makeFooRight() (Foo, error) {
	f := Foo{
		foo: "Hello",
		bar: 1,
	}

	return f, nil
}

func jsonExample() {
	f := struct {
		Foo string `json:"foo"` // Capital property name on the struct and no spaces between
		Bar int    `json:"bar"`
	}{}
	fmt.Println("f before json.Unmarshal:", f)

	// With json functions we pass pointers because it did not support generics.
	// Also to optimize memory allocation and garbage collection in case Unmarshal is called in a loop.
	err := json.Unmarshal([]byte(`{ "foo": "fooJson", "bar": 2}`), &f)

	if err != nil {
		fmt.Println("Error while parsing the json:", err)
	}
	fmt.Println("f after json.Unmarshal:", f)

	f2 := FooJson{}
	fmt.Println("f2 before json.Unmarshal:", f2)

	err = json.Unmarshal([]byte(`{ "foo": "fooJson", "bar": 2}`), &f2)

	if err != nil {
		fmt.Println("Error while parsing the json:", err)
	}
	fmt.Println("f after json.Unmarshal:", f2)
}

func funcWithLargeInput(val struct {
	arr []int
}) {
	val.arr[10] = 1
}

func funcWithLargeInputPointer(val *struct {
	arr []int
}) {
	val.arr[10] = 2
}

func withLargeOutput() struct {
	arr []int
} {
	largeArray := make([]int, 20_000)
	objectWithLargeArray := struct {
		arr []int
	}{
		arr: largeArray,
	}

	return objectWithLargeArray
}

func withLargeOutputPointer() *struct {
	arr []int
} {
	largeArray := make([]int, 20_000)
	objectWithLargeArray := struct {
		arr []int
	}{
		arr: largeArray,
	}

	return &objectWithLargeArray
}

func returnNoValueWrong() *int {
	return nil
}

func returnNoValueCorrect() (int, bool) {
	return 0, false
}

func pointersUseCases() {
	var f1 *Foo
	makeFooWrong(f1)
	f2, _ := makeFooRight()
	fmt.Println("f1:", f1)
	fmt.Println("f2:", f2)
	fmt.Println()

	jsonExample()
	fmt.Println()

	largeArray := make([]int, 20_000)
	objectWithLargeArray := struct {
		arr []int
	}{
		arr: largeArray,
	}
	beforeTime := time.Now()
	funcWithLargeInput(objectWithLargeArray)
	afterTime := time.Now()
	fmt.Println("Time difference input struct:", afterTime.UnixNano()-beforeTime.UnixNano()) // with 20_000 size array it already consistently returns a difference.

	beforeTime = time.Now()
	funcWithLargeInputPointer(&objectWithLargeArray)
	afterTime = time.Now()
	fmt.Println("Time difference input pointer:", afterTime.UnixNano()-beforeTime.UnixNano()) // This rarely returns anything above 0

	// Returning pointers is different though.
	beforeTime = time.Now()
	withLargeOutput()
	afterTime = time.Now()
	fmt.Println("Time difference return struct:", afterTime.UnixNano()-beforeTime.UnixNano())

	beforeTime = time.Now()
	withLargeOutputPointer()
	afterTime = time.Now()
	fmt.Println("Time difference return pointer:", afterTime.UnixNano()-beforeTime.UnixNano())
	fmt.Println()

	x := returnNoValueWrong()
	if x == nil {
		fmt.Println("You should not use a nil pointer as an undefined value", x)
	}

	y, ok := returnNoValueCorrect()
	if ok == false {
		fmt.Println("No value has been returned by the function:", y, ok)
	}
}

func pointersExamples() {
	pointerSyntaxExamples()
	fmt.Println()

	pointerTypeExamples()
	fmt.Println()

	reassingingPointersExamples()
	fmt.Println()

	pointersUseCases()
}

func main() {
	display.SectionTitle("Pointers")
	pointersExamples()
}
