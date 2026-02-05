package main

import "fmt"

func main() {
	// Literal types
	// integerLiteral, _ is used for readability
	integerLiteral := 1_234
	fmt.Println(integerLiteral)
	// floatLiteral
	floatLiteral := 2.345e2
	fmt.Println(floatLiteral)
	// hexadecimalLiteral, p means power of 2
	hexadecimalLiteral := 0x12.34p0
	fmt.Println(hexadecimalLiteral)
	// 0 prefix means octalLiteral
	octalLiteral := 012
	fmt.Println(octalLiteral)
	// characted
	charLiteral := 'a'
	fmt.Printf("char: %c\n", charLiteral)
	// char 8 bit octal
	charOctalLiteral := '\141'
	fmt.Printf("char 8 bit octal: %c\n", charOctalLiteral)
	// char 8 bit hexadecimal
	charHexLiteral := '\x61'
	fmt.Printf("char 8 bit hexadecimal: %c\n", charHexLiteral)
	// char 16 bit hexadecimal
	charHex16Literal := '\u0061'
	fmt.Printf("char 16 bit hexadecimal: %c\n", charHex16Literal)
	// char 32 bit unicode
	charUnicodeLiteral := '\U00000061'
	fmt.Printf("char 32 bit unicode: %c\n", charUnicodeLiteral)
	stringLiteral := "Greetings and\n\"Salutations!\""
	fmt.Printf("string literal: %s\n", stringLiteral)
	// raw string literal
	rawStringLiteral := `Greetings and
"Salutations!"`
	fmt.Printf("raw string: %s\n", rawStringLiteral)

	// Boolean
	var flag bool // zero value is false
	fmt.Printf("flag: %t\n", flag)
	var trueValue bool = true
	fmt.Printf("trueValue: %t\n", trueValue)

	// Numeric Types
	// Integer Types
	var integer int                      // int depends on architecture running the program and must be casted to be compared to other int types.
	fmt.Printf("integer: %d\n", integer) // Defaults to 0
	var integer8 int8 = 127
	fmt.Printf("integer8: %d\n", integer8)
	var unsigned8 uint8 = 20           // int can be 8, 16, 32 and 64 bytes
	var byteValue byte = unsigned8 + 8 // alias for uint8
	fmt.Printf("byteValue: %d\n", byteValue)
	var integer64 int64 = -9223372036854775808
	fmt.Printf("integer16: %d\n", integer64)
	// uint depends on architecture running the program. Can't be compared to int
	var unsignedInteger uint = 123456
	fmt.Printf("unsignedInteger: %d\n", unsignedInteger)
	// Integer operations
	var a int64 = 24
	var b int64 = 37
	// var c int32 = 9
	// var d = b + c // Invalid due to type mismatch
	fmt.Printf("a: %d; b: %d\n", a, b)
	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b) // using uint for the same operation resets the count from big numbers
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %d\n", a/b)
	fmt.Printf("b modulus a = %d\n", b%a)
	fmt.Printf("bitwise shift left %d\n", a<<1)
	fmt.Printf("bitwise shift right %d\n", a>>1)
	fmt.Printf("a bitwise and b: %d\n", a&b)
	fmt.Printf("a bitwise or b: %d\n", a|b)
	fmt.Printf("a bitwise xor b: %d\n", a^b)
	fmt.Printf("a bitwise and not b: %d\n", a&^b)
}
