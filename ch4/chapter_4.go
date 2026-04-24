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

	display.SectionTitle("If block")
	ifBlockExamples()
}
