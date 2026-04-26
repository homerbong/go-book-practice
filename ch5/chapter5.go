package main

import (
	"errors"
	"fmt"
	"go-book-practice/display"
)

type MyFuncParams struct {
	FirstName string
	LastName  string
	Age       int
}

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
}

func main() {
	display.SectionTitle("Functions")
	functionsExamples()
}
