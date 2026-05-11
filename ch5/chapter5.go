package main

import (
	"errors"
	"fmt"
	"go-book-practice/display"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type MyFuncParams struct {
	FirstName string
	LastName  string
	Age       int
}

/*
Type for operator functions
*/
type opFuncType func(int, int) int

type person struct {
	age  int
	name string
}

// Use with very much care. You shouldn't do this unless you really need this capability.
// Package level state should be immutable to make data flow easier to understand.
var (
	add = func(i, j int) int { return i + j }
	sub = func(i, j int) int { return i - j }
	mul = func(i, j int) int { return i * j }
	// div = func(i, j int) int { return i / j }
)

func div(num, denom int) int { // You can use only one type at the end if all params are of the same type.
	if denom == 0 {
		return 0
	}

	return num / denom
}

func MyFunc(params MyFuncParams) {
	fmt.Println("First name: ", params.FirstName)
	fmt.Println("Last  name: ", params.LastName)
	fmt.Println("Age       : ", params.Age)
}

func addTo(base int, valsToAdd ...int) []int {
	out := make([]int, 0, cap(valsToAdd))

	for _, val := range valsToAdd {
		out = append(out, val+base)
	}

	return out
}

func divAndRemainder(num int, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("Cannot divide by 0")
	}

	return num / denom, num % denom, nil
}

func divAndRemainderNamedReturns(num int, denom int) (result int, remainder int, err error) {
	result, remainder = 10, 20
	if denom == 0 {
		err = errors.New("Cannot divide by 0") // Predeclared variables. They can be shadowed
		return 0, 0, err                       // You can still return values but Go will add code to assign the values in the return statement to the return variables.
	}

	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func divAndRemainderBlankReturn(num int, denom int) (result int, remainder int, err error) {
	result = 10
	if denom == 0 {
		err = errors.New("Cannot divide by 0")
		return // Never use this: it makes confusing trying to understand where the values were assigned.
	}

	result, remainder = num/denom, num%denom
	return
}

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}

	return total
}

func calculatorExample() {
	add := func(i, j int) int { return i + j }
	sub := func(i, j int) int { return i - j }
	mul := func(i, j int) int { return i * j }
	div := func(i, j int) int { return i / j }
	var opMap = map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	var expressions = [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, val := range expressions {
		if len(val) != 3 {
			fmt.Println("Invalid expression:", val)
			continue
		}

		p1, err := strconv.Atoi(val[0])
		if err != nil {
			fmt.Println("Error parsing first operand:", err)
			continue
		}

		opFunc, ok := opMap[val[1]]
		if !ok {
			fmt.Println("Invalid operator:", val[1])
			continue
		}

		p2, err := strconv.Atoi(val[2])
		if err != nil {
			fmt.Println("Error parsing second operand", err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(val, "=", result)
	}
}

func anonyousFunctionsExample() {
	// Declared anonymous function
	f := func(j int) {
		fmt.Println("Printing", j, "from an anonymous function")
	}

	for i := range 5 {
		f(i)
	}
	fmt.Println()

	for i := range 5 {
		// Immediately called function: not really useful unless you're using the defer statement or launching goroutines
		func(j int) {
			fmt.Println("Printing", j, "from an anonymous function")
		}(i)
	}
	fmt.Println()

	result := add(1, 5)
	fmt.Println("original -> 1 + 5 =", result)
	// You can reassing an anonymous function declared as package varriable.
	add = func(i, j int) int {
		return i + j + j
	}
	result = add(1, 5)
	fmt.Println("new      -> 1 + 5 =", result)
	fmt.Println()

	result = div(4, 2)
	fmt.Println("original -> 4 / 2 =", result)
	// You can shadow a package function but you can't reassign it!
	// div = func(i, j int) int { // compiler error: cannot assign to div (neither addressable nor a map index expression
	// 	return i / (2 * j)
	// }
	div := func(i, j int) int { // Changes how div works but it only does for this scope.
		return i / (2 * j)
	}
	result = div(4, 2)
	fmt.Println("new      -> 4 / 2 =", result)
}

func makeMult(base int) func(int) int {
	return func(val int) int {
		return base * val
	}
}

func closuresExamples() {
	a := 20
	f := func() {
		fmt.Println("inside f  -> a =", a)
		a += 10
	}
	f()
	fmt.Println("outside f -> a =", a)
	f()
	fmt.Println("outside f -> a =", a)
	fmt.Println()

	// You can shadow a variable in the closure
	f2 := func() {
		a := 100
		fmt.Println("inside f2  -> a =", a)
	}
	f2()
	fmt.Println("outside f2 -> a =", a)
	fmt.Println()

	// Why closures?
	// It is useful to limit a function's scope: you use it multiple times here but not in other places.
	// Also useful when you need to work on a variable declared in the same function and access from your closures.
	// Also useful when you return a function from another function as it allows modifying the variables of the containing function.

	// Passing functions as parameters:
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println("people before sorting:            ", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("people after sorting by last name:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("people after sorting by age      :", people)
	fmt.Println()

	// Just to understand the algorithm behind sort.
	numbers := []int{4, 6, 5, 2, 7}
	sort.Slice(numbers, func(i, j int) bool {
		fmt.Println(i, j)
		return numbers[i] < numbers[j]
	})
	fmt.Println()
	// Returning functions

	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := range 3 {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

func cat() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func multipleDeferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first defer -> a:", val)
	}(a)

	a = 20
	defer func(val int) {
		fmt.Println("second defer -> a:", val)
	}(a)

	a = 30
	fmt.Println("exiting a: ", a)
	return a
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, nil, err
	}

	return file, func() {
		file.Close()
	}, err
}

func simpleGetFileExample() {
	file, closer, err := getFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer closer()

	stat, err := file.Stat()
	fmt.Println("File size:", stat.Size())
}

func deferExamples() {
	cat()

	fmt.Println()
	a := multipleDeferExample()
	fmt.Println("returned a:", a)

	fmt.Println()
	simpleGetFileExample()
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Sam"
	fmt.Println("Inside modify(i, s, p):", i, s, p)
}

func modMap(m map[int]string) {
	m[2] = "Hello"
	m[3] = "Goodbye"
	delete(m, 1)
	fmt.Println("Inside modMap:", m)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
	fmt.Println("Inside modSlice:", s)
}

func callByValueExamples() {
	p := person{}
	i := 1
	s := "Hello"

	fmt.Println("Before modifyFails(i, s, p):", i, s, p)
	modifyFails(i, s, p)
	fmt.Println("Outside modify(i, s, p):", i, s, p)
	fmt.Println()

	m := map[int]string{
		1: "first",
		2: "second",
	}
	fmt.Println("Before modMap:", m)
	modMap(m)
	fmt.Println("Outside modMap:", m)
	fmt.Println()

	sl := []int{3, 4, 5}
	fmt.Println("Before modSlice:", sl)
	modSlice(sl)
	fmt.Println("Outside modSlice:", sl)
}

func functionsExamples() {
	result := div(5, 2)
	fmt.Println("result:", result)
	fmt.Println()

	fmt.Println("Optional params")
	MyFunc(MyFuncParams{LastName: "Doe", Age: 24})
	fmt.Println()
	MyFunc(MyFuncParams{FirstName: "John", LastName: "Doe"})
	fmt.Println()

	fmt.Println("Variadic Input Parameters and Slices")
	fmt.Println("1 param call    :", addTo(1))
	fmt.Println("2 param call    :", addTo(2, 3))
	fmt.Println("n param call    :", addTo(2, 2, 5, 6))
	a := []int{4, 7, 2, 7}
	fmt.Println("slice param call:", addTo(2, a...))
	fmt.Println("slice param call:", addTo(2, []int{3, 6, 4, 7, 9}...))
	fmt.Println()

	fmt.Println("Multiple Return Values")
	result, remainder, err := divAndRemainder(5, 3)
	fmt.Println("result, remainder, error:", result, remainder, err)
	// Ignoring values
	_, _, err = divAndRemainder(3, 0)
	if err != nil {
		fmt.Println("error:", err)
	}
	// Named return values
	x, y, z := divAndRemainderNamedReturns(7, 2)
	fmt.Println("result, remainder, error:", x, y, z)
	result, remainder, err = divAndRemainderBlankReturn(5, 0)
	fmt.Println("result, remainder, error:", result, remainder, err) // Never use this

	display.SectionTitle("Functions are values")
	var myFuncVar func(string) int // default value is nil
	// fmt.Println("myFunc:", myFunc) // vet error: fmt.Println arg myFunc is a func value, not called
	myFuncVar = f1
	fmt.Println("Calling myFuncVar(\"Hello`\"): ", myFuncVar("Hello"))
	myFuncVar = f2
	fmt.Println("Calling myFuncVar(\"Hello`\"): ", myFuncVar("Hello"))
	calculatorExample()

	display.SectionTitle("Anonymous functions")
	anonyousFunctionsExample()
	fmt.Println("add after being modified by other function -> 2 + 3 = ", add(2, 3))

	display.SectionTitle("Closures")
	closuresExamples()

	display.SectionTitle("defer")
	deferExamples()

	display.SectionTitle("Go is call by value")
	callByValueExamples()
}

func exercise51() {
	add, sub, mul, div, rem := func(a, b int) (int, error) {
		return a + b, nil
	}, func(a, b int) (int, error) {
		return a - b, nil
	}, func(a, b int) (int, error) {
		return a * b, nil
	}, func(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("Can't divide by 0!")
		}
		return a / b, nil
	}, func(a, b int) (int, error) {
		return a % b, nil
	}

	opFunc := map[string]func(int, int) (int, error){
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
		"%": rem,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"3", "%", "2"},
		{"4", "/", "0"},
		{"two", "+", "three"},
		{"5"},
	}

	fmt.Println("Exercise 1")
	fmt.Println()

	for _, v := range expressions {
		if len(v) != 3 {
			fmt.Println("Invalid length of the expressions")
			continue
		}

		firstOperand, err := strconv.Atoi(v[0])
		if err != nil {
			fmt.Println("Invalid first operand:", v[0])
			continue
		}

		operator := v[1]
		operation := opFunc[operator]
		if operation == nil {
			fmt.Println("Invalid Operation: ", operator)
			continue
		}

		secondOperand, err := strconv.Atoi(v[2])
		if err != nil {
			fmt.Println("Invalid second operand:", v[2])
		}

		res, err := operation(firstOperand, secondOperand)
		if err != nil {
			fmt.Println("The operation failed:", err)
			continue
		}
		fmt.Println(v, "=", res)
	}
}

func fileLenSolution(name string) (int, error) {
	file, err := os.Open(name)

	if err != nil {
		return 0, err
	}

	defer func() {
		file.Close()
	}()

	bytes := make([]byte, 2048)

	totalLen := 0
	for {
		count, err := file.Read(bytes)

		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			return totalLen, nil
		}

		totalLen += count
	}
}

func fileLen(name string) (int, error) {
	file, err := os.Open(name)

	if err != nil {
		return 0, err
	}

	defer func() {
		file.Close()
	}()

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return int(stat.Size()), nil
}

func exercise52() {
	fmt.Println("Exercise 2")
	if len(os.Args) < 2 {
		fmt.Println("File path not provided")
		return
	}

	l, err := fileLen(os.Args[1])
	if err != nil {
		fmt.Println("Error while obtaining the length of the file:", err)
	}

	fmt.Println("Length of the file:", l)

	l, err = fileLenSolution(os.Args[1])
	if err != nil {
		fmt.Println("Error while obtaining the length of the file:", err)
	}

	fmt.Println("Length of the file:", l)
}

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}

func exercise53() {
	fmt.Println("Exercise 3")
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Sam"))
}

func exercises() {
	display.SectionTitle("Exercises")
	fmt.Println()
	exercise51()

	fmt.Println()
	exercise52()

	fmt.Println()
	exercise53()
}

func main() {
	display.SectionTitle("Functions")
	functionsExamples()
	exercises()
}
