package main

// Each place where a declaration occurs is called a block.

// This is the import block
import (
	"fmt"
	"go-book-practice/display"
	"math/rand"
)

// The import block can contain variables, constants, types and functions.
const packageConstant = "Package Constant"

var packageVar = "Package Variable"

func example42() {
	fmt.Println()
	x := 10
	fmt.Println("Another example of shadowing the variables x =", x)
	if x > 5 {
		x, y := 5, 20 // shadowing the variable again
		fmt.Println("Inside the if block: x, y  =", x, y)
	}
	fmt.Println("Again x = ", x)
}

/*
Idiomatic way of writing the FizzBuzz logic.
Avoid nesting ifs. Use continue to make code clearer
*/
func example412() {
	fmt.Println("FizzBuzz example")
	for i := 1; i <= 50; i++ {
		if i != 1 {
			fmt.Print(", ")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Print("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
			continue
		}
		fmt.Print(i)
	}

	fmt.Println()
}

func example415() {
	fmt.Println("The order of keys is not guaranteed in Go")
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

func example416() {
	fmt.Println("Strings iterate over runes, not over bytes. The first value indicates the number of bytes from the start though.")
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

func shadowingPackages() {
	fmt.Println()
	fmt.Println("Shadowing packages:")
	x := 10
	fmt.Println("Printing still works: ", x)
	fmt.Println("Now fmt will be shadowed and the next print of x won't be appearing")
	fmt := struct {
		Println func(param string)
	}{
		Println: func(param string) {
			// Does nothing anymore
		},
	}
	fmt.Println(string(rune(x)))
}

func universeBlockShadowing() {
	fmt.Println()
	fmt.Println("Value of true:", true)
	true := "true changed to string"
	fmt.Println("Shadowed value of true: ", true)
}

func ifBlockExamples() {
	n := rand.Intn(10)

	if n == 0 {
		fmt.Println("Random number is equal to 0: ", n)
	} else if n > 5 {
		fmt.Println("Random number is bigger than 5: ", n)
	} else {
		fmt.Println("Random number is between 1 and 5: ", n)
	}

	// Scoping variables just for if and else blocks
	if n := rand.Intn(20); n == 0 {
		fmt.Println("if block variable is equal to 0: ", n)
	} else if n > 10 {
		fmt.Println("if block random number is bigger than 10: ", n)
	} else {
		fmt.Println("if block random number is between 1 and 10: ", n)
	}
	fmt.Println("n has been shadowed and its values is: n = ", n)

	if x := rand.Int(); x > -1 { // Technically you can put any simple statement before the comparison. DON'T DO IT! Just declare variables you need for the if else statement
		fmt.Println("This will always be true: x = ", x)
	}
	// fmt.Println(x) // compiler error: undefined: x
}

func forBlockExamples() {
	// complete for statement
	fmt.Println("Complete for statement: ")
	fmt.Print("i -> ")
	for i := 0; i < 10; i += 1 {
		fmt.Print(i)
		if i < 9 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
	fmt.Println()

	idx := 0
	fmt.Println("for statement without initialization")
	fmt.Print("i -> ")
	for ; idx < 10; idx += 2 {
		fmt.Print(idx)
		if idx < 8 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
	fmt.Println("idx at the end: ", idx) // idx is being declared outside and it's being modified
	fmt.Println()

	fmt.Println("for statement without increment")
	fmt.Print("i -> ")
	for i := 0; i < 10; {
		fmt.Print(i)
		if i == 0 {
			i++
		} else {
			i += 2
		}

		if i <= 9 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
	fmt.Println()

	// The condition only for statement
	idx = 0
	fmt.Println("Condition only for statement")
	fmt.Print("i -> ")
	for idx < 10 { // no semicolons needed
		fmt.Print(idx)
		if idx < 8 {
			fmt.Print(", ")
		}
		idx += 2
	}
	fmt.Println()
	fmt.Println()

	// The infinite for loop
	idx = 1
	fmt.Println("The infinite for statement")
	for {
		if n := rand.Intn(20); n == 18 {
			fmt.Println()
			fmt.Printf("Interrupting the for loop because the random number is 18 after %d cycles", idx)
			break
		}
		idx++
	}
	fmt.Println()

	example412()
	fmt.Println()

	// The for range statement
	fmt.Println("The for range statement")
	evenVals := []int{2, 4, 6, 8, 10}
	indexes := ""
	values := ""
	for i, v := range evenVals {
		indexes = indexes + string(fmt.Sprint(i))
		values = values + string(fmt.Sprint(v))
		if i < len(evenVals)-1 {
			indexes = indexes + ", "
			values = values + ", "
		}
	}
	fmt.Println("indexes: ", indexes)
	fmt.Println("values : ", values)

	wins := map[string]int{"Fred": 1, "Mike": 4, "Alex": 2}
	for k, v := range wins {
		fmt.Println(k, ": ", v)
	}

	values = ""
	for _, v := range evenVals { // You can ignore the first variable with underscore
		values = values + string(fmt.Sprint(v))
		if v != evenVals[len(evenVals)-1] {
			values = values + ", "
		}
	}
	fmt.Println("values: ", values)

	keys := make([]string, 0, 10)
	for k := range wins {
		keys = append(keys, k)
	}
	fmt.Println("keys: ", keys)

	fmt.Println()
	example415()
	fmt.Println()

	fmt.Println("Iterating over strings")
	example416()

	fmt.Println("For values are a copy")
	for _, v := range evenVals { // This is created on each iteration. Useful for go routines.
		v *= 2
	}
	fmt.Println("Unchanged evenVals: ", evenVals)

	// labels
	fmt.Println("Labels")

	samples := []string{"hello", "apple_π!"}
outer:
	for _, v := range samples {
		fmt.Println(v)
	inner:
		for _, r := range v {
			if r == 'e' || r == 'p' {
				continue inner
			}

			if r == 'l' {
				continue outer
			}
			fmt.Println(string(r))
		}
		fmt.Println()
	}
}

func compareForLoops() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		if i == 0 {
			continue
		}
		if i == len(evenVals)-1 {
			break
		}
		fmt.Println(i, v)
	}

	// VS

	for i := 1; i < len(evenVals)-1; i++ { // not working for strings
		fmt.Println(i, evenVals[i])
	}
}

func switchStatementExamples() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size { // statement scoped variable
		case 1, 2:
			fallthrough // It allows for this kind of logic but you should reconsider using it.
		case 3, 4:
			fmt.Println(word, "is a short word")
			if word != "cow" { // Mulitple lines under the same case without braces
				break // It allows for the break statement but it might indicate you're doing something too complicated.
			}
			fmt.Println("Muuuuuuuu")
		case 5:
			wordLen := len(word)                                       // Case scoped variable
			fmt.Println(word, "is exactly the right length:", wordLen) // No break statement
		case 6, 7, 8, 9: // Nothing happens for these cases
		default:
			fmt.Println(word, "is a long word!")
		}
	}
	fmt.Println()

loop:
	for i := 0; i < 10; i += 1 {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not by 2")
		case 7:
			fmt.Println("Exit the loop") // 7, 8 and 9 won't be printed.
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
	fmt.Println()

	// blank switch
	otherWords := []string{"hi", "salutations", "hello"}
	for _, word := range otherWords {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length:", wordLen)

		}
	}

	// regular switch vs blank switch
	// if you find yourself comparing the same variable in the case statements such as below
	a := 1
	switch {
	case a == 1:
		fmt.Println("a is 1")
	case a == 2:
		fmt.Println("a is 2")
	case a == 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is", a)
	}
	fmt.Println()
	// You should probably rewrite this as
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is", a)
	}
	fmt.Println()

	// Choosing between if and switch
	// A switch statement indicates there is a relationship between the values and comparisons in each case
	// FizzBuzz example
	for i := 1; i < 31; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i) // No need for continue statements and default behaviour is explicit.
		}
	}
}

/*
Not recommended to use unless there are particular scenarios that require it.
There are limitations to help avoiding misuse.
*/
func gotoStatementExamples() {
	a := 10
	// goto skip // vet error: goto skip jumps over variable declaration at line 399
	b := 20

	// skip:
	c := 30
	fmt.Println(a, b, c)
	if c > a {
		fmt.Println("c is greater than a")
		// goto inner // compiler error:  goto inner jumps into block starting at ./chapter_4.go:408:11
	}
	if a < b {
		// inner:
		fmt.Println("a is less than b")
	}
	fmt.Println()

	// Legal use of the goto statement
	a = rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done // This could be done with a boolean flag or duplicate the code after.
			// boolean is the same as goto but more verbose.
			// Duplicating code makes it harder to maintain
		}
		a = a*2 + 1
	}
	fmt.Println("Do something if the for loop completes.")
done:
	fmt.Println("Do complicated stuff no matter why we left the loop")
	fmt.Println("Done: ", a)

	fmt.Println()
	fmt.Println("A very useful example on where goto is useful can be found in the floatBits method in the file atof.go in the strconv package")
}

func exercise1() []int {
	randNums := make([]int, 0, 100)
	for range 100 {
		randNums = append(randNums, rand.Intn(101))
	}
	fmt.Println("randNums: ", len(randNums), cap(randNums))
	fmt.Println(randNums)
	return randNums
}

func exercise2(nums []int) {
	for _, val := range nums {
		switch {
		case val%2 == 0 && val%3 == 0:
			fmt.Println("Six!")
		case val%2 == 0:
			fmt.Println("Two!")
		case val%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}

func exercise3() {
	var total int
	for i := range 10 {
		total := total + i // Shadowing
		fmt.Println("total in loop:", total)
	}
	fmt.Println("total outside:", total) // 0 as it was shadowed by the for block declared variable.
}

func main() {
	display.SectionTitle("Chapter 4")
	fmt.Println("hello")
	{
		// This is a block within the main function's block.
		fmt.Println("This is within a block")
	}
	display.SectionTitle("Shadowing Variables")
	fmt.Println("packageVar in main before assignment: ", packageVar)
	packageVar := "Package Variable in main" // := reuses only variables declared in the same block.
	fmt.Println("packageVar in main after assignment: ", packageVar)
	{

		fmt.Println("packageVar within the block before assignment: ", packageVar)
		packageVar := "changed value" // The var declared in the package block is being shadowed.
		fmt.Println("packageVar within the block: ", packageVar)
	}
	fmt.Println("packageVar outside the block again: ", packageVar)
	example42()

	shadowingPackages()
	fmt.Println("Now Println works again")

	universeBlockShadowing()

	display.SectionTitle("if Statement")
	ifBlockExamples()

	display.SectionTitle("for Statement")
	forBlockExamples()
	// Use for range statement with built-in compound types as it makes thigs easier
	// Use complete for statement when iterating over a subset of a compound type. (slices could also work?)
	// Condition only is a good replacement for the while loop
	// infinite for loop might be usefule in very rare cases and it should have a break or return statement. Also useful for the iterator pattern.

	display.SectionTitle("switch Statement")
	switchStatementExamples()

	display.SectionTitle("goto Statement")
	gotoStatementExamples()

	display.SectionTitle("Exercise 1")
	randNums := exercise1()

	display.SectionTitle("Exercise 2")
	exercise2(randNums)

	display.SectionTitle("Exercise 3")
	exercise3()
}
