package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-book-practice/display"
	"math/rand"
	"os"
	"runtime"
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

type Person struct {
	firstName string
	lastName  string
	age       int
}

type A struct {
	b *B
}

type B struct {
	c *C
}

type C struct {
	field string
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

/*
Map should be used as an input param only if keys are not known before hand.
*/
func mapInputParameter(m map[string]int) {
	for k, v := range m {
		// Process unknown keys and their values.
		fmt.Println("k, v:", k, v)
	}
}

/*
Structs should be used to define complex objects instead of generic maps.
*/
func structInputParam(data struct {
	firstKey  string
	secondKey int
}) {
	fmt.Println("firstKey:", data.firstKey)
	fmt.Println("secondKey:", data.secondKey)
}

/*
Slices are passed as structs containing len, cap and a pointer in memory.
Changing elements in the copied struct pointer changes it in the original slice.
*/
func changingElements(sl []int) {
	for k, v := range sl {
		sl[k] = v + 1*k
	}

	fmt.Println("inside changingElements")
	fmt.Println("sl, len, cap:", sl, len(sl), cap(sl))
}

/*
Chaging the length of a slice does not affect the original slice because the
value of len in the original struct is not being updated.
*/
func changingLenDoesNotReflect(sl []int) {
	sl = append(sl, 5)
	fmt.Println("Inside changingLenDoesNotReflect")
	fmt.Println("sl, len, cap:", sl, len(sl), cap(sl))
}

func appendingToSlice(sl []int) {
	sl = append(sl, 6)
	fmt.Println("inside appending:")
	fmt.Println("sl, len, cap:", sl, len(sl), cap(sl))
}

/*
You can modify the content of a slice but not it's length or cap.
It is useful when using as a buffer for data to be read or written.
*/
func randomBytesBuffer(buf []byte) []byte {
	bufLen := len(buf)

	if bufLen > 0 && bufLen%2 != 0 {
		fmt.Println("An even length buffer must be provided for UTF-8 encoding to be correct.")
		return buf
	}

	for i := 1; i < bufLen; i += 2 {
		randVal := rand.Intn(95)
		randVal += 32
		buf[i] = byte(randVal)
		buf[i-1] = byte(0)
	}

	return buf
}

/*
To demonstrate the use of slices as buffers we
*/
func createFileWithRandomData() {
	file, err := os.Create("example")

	if err != nil {
		fmt.Println("Error while creating the file:", err)
	}

	defer func() {
		fmt.Println("Created file with random data")
		file.Close()
	}()

	buf := make([]byte, 256)
	newLine := []byte{byte('\n')}

	for range rand.Intn(32) {
		buf = randomBytesBuffer(buf)

		_, err := file.Write(buf)
		if err != nil {
			fmt.Println("Error while writing the buffer to the file.")
		}
		file.Write(newLine)
	}
}

/*
The advantage of passing a slice as a param is the fact that it passes a struct (pointer, len, cap)
When passing an array the whole array is being passed.
*/
func slicePointersExamples() {
	sl := make([]int, 3, 6)

	fmt.Println("before changingElements")
	fmt.Println("sl, len, cap", sl, len(sl), cap(sl))
	changingElements(sl)
	fmt.Println("after changingElements")
	fmt.Println("sl, len, cap", sl, len(sl), cap(sl))
	fmt.Println()

	fmt.Println("Before changingLenDoesNotReflect")
	fmt.Println("sl, len, cap:", sl, len(sl), cap(sl))
	changingLenDoesNotReflect(sl)
	fmt.Println("after changingLenDoesNotReflect")
	fmt.Println("sl, len, cap:", sl, len(sl), cap(sl))
	// fmt.Println("sl[4]:", sl[3]) // panic: runtime error: index out of range [3] with length 3
	fmt.Println()

	sl = []int{1, 2, 3}
	fmt.Println("Before appending")
	fmt.Println("sl, len, cap:", sl, len(sl), cap(sl))
	appendingToSlice(sl)
	fmt.Println("Before appending")
	fmt.Println("sl, len, cap:", sl, len(sl), cap(sl))
	fmt.Println()

	createFileWithRandomData()
}

func pointersToMapsAndSlicesExamples() {
	m := make(map[string]int)
	m["A"] = 30
	m["B"] = 87

	mapInputParameter(m)

	obj := struct {
		firstKey  string
		secondKey int
	}{
		firstKey:  "Hello",
		secondKey: 54,
	}

	structInputParam(obj)
	fmt.Println()

	slicePointersExamples()
}

func makeAPointer() *A {
	a := &A{&B{&C{"Hello"}}}

	runtime.SetFinalizer(a.b.c, func(c *C) { fmt.Println("a.b.c with value", c.field, "is garbage collected") })

	return a
}

func garbageCollectionExample() {
	aPointer := makeAPointer()

	runtime.GC()

	time.Sleep(200)
	fmt.Println("aPointer: ", aPointer)

	aPointer = nil
	fmt.Println("aPointer: ", aPointer)

	runtime.GC()

	time.Sleep(200)

	fmt.Println("Exiting garbageCollectionExample")
}

func pointersExamples() {
	pointerSyntaxExamples()
	fmt.Println()

	pointerTypeExamples()
	fmt.Println()

	reassingingPointersExamples()
	fmt.Println()

	pointersUseCases()
	fmt.Println()

	pointersToMapsAndSlicesExamples()

	garbageCollectionExample()
}

func MakePerson(firstName string, lastName string, age int) Person {
	// The return value escapes the heap because a struct is a pointer.
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	// It does escape the heap because it's still returning a pointer.
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func exercise61() {
	newPerson := MakePerson("Dan", "Peterson", 30)
	// newPerson escapes to heap because the current Go compiler
	// moves to the heap any value that is passed in to a function
	// via a parameter that is of an interface type.
	fmt.Println("newPerson:", newPerson)

	newPersonPointer := MakePersonPointer("Dan", "Peterson", 30)
	newPersonPointer.age = 31

	// Does not escape the heap because is the same pointer.
	fmt.Println("newPersonPointer:", newPersonPointer)

	// It escapes the heap because it's a value being passed
	fmt.Println("newPersonPointer:", *newPersonPointer)
}

func exercises() {
	display.SectionTitle("Exercises")

	exercise61()
}

func main() {
	display.SectionTitle("Pointers")
	pointersExamples()

	exercises()
}
