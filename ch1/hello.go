package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, World!")
// }

// Exercise 3
func main() {
	// Can't end with new line
	fmt.Println("Hello, World! \n\bAnother Line")
	fmt.Println("	Second print	\bNow what?	Asd")
	// fmt.Println("something\n") // This will cause a vet error but not a compiler error
}

// testing go vet

// func main() {
// 	fmt.Printf("Hello %s,\n")
// }
