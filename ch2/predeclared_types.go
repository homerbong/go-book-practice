package main

import (
	"fmt"
	"math/cmplx"
)

func shadowingExample() {
	var x = 10
	fmt.Println("x (outside of the if) = ", x)
	if x > 5 {
		fmt.Println("x (inside of the if shaded) = ", x)     // Takes the value of the above declared var
		x := 5                                               // New shadowing variable is created
		fmt.Println("x (outside of the if shadowing) = ", x) // Shows the shadowing value
	}
	fmt.Println("x (outside of the if shaded) = ", x) // Shaded value again
}

func exercise1() {
	var i int = 20
	var f float64 = float64(i)

	fmt.Println("exercise 1 i, f:", i, f)
}
func exercise2() {
	const value = 7

	var i int = value
	var f float64 = value

	fmt.Println("exercise 2 i, f:", i, f)
}

func exercise3() {
	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI int64 = 9_223_372_036_854_775_807

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println("exercise 3 b, smallI, bigI:", b, smallI, bigI)
}

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

	// Float types
	var floatA float32 = 3.2
	var floatB float32 = 4.5
	fmt.Printf("floatA + floatB = %f\n", floatA+floatB)
	fmt.Printf("floatA + floatB = %f\n", floatA-floatB)
	fmt.Printf("floatA + floatB = %f\n", floatA*floatB)
	fmt.Printf("floatA + floatB = %f\n", floatA/floatB)
	// All integer operators are also valid for floats

	var cFloat = floatA + floatB
	var eps float32 = 0.000001
	fmt.Printf("cFloat == 7.7 -> %t\n", cFloat == 7.7)         // Not recommended
	fmt.Printf("cFloat - 7.7 < eps -> %t\n", cFloat-7.7 < eps) // Recommended way of using floats.

	// Complex numbers
	var complexA complex128 = complex(20.3, 10.2)
	var complexB complex128 = complex(1.8, 8.2)
	fmt.Println("complexA + complexB = ", complexA+complexB)
	fmt.Println("complexA - complexB = ", complexA-complexB)
	fmt.Println("complexA * complexB = ", complexA*complexB)
	fmt.Println("complexA / complexB = ", complexA/complexB)
	fmt.Println("real part of complexA ->", real(complexA))
	fmt.Println("imaginary part of complexA ->", imag(complexA))
	fmt.Println("Modulus of complexA -> ", cmplx.Abs(complexA))

	// Runes
	var myFirstInitial rune = 'A' // good - the type name matches the usage
	var myLastInitial int32 = 'M' // bad - legal but confusing
	fmt.Println("myFirstInitial: ", myFirstInitial)
	fmt.Printf("myFirstInitial %c\n", myFirstInitial)
	fmt.Println("myLastInitial: ", myLastInitial)
	fmt.Printf("myLastInitial %c\n", myLastInitial)

	// Explicit type conversion
	var anInt int = 10
	var aFloat float64 = 30.2
	var aByte byte = 100
	// var anotherByte byte = 270 Illegal with vet.
	var mixingSum1 = float64(anInt) + aFloat
	var mixingSum2 = anInt + int(aFloat)
	var mixingSum3 = anInt + int(aByte)
	var mixingSum4 = byte(anInt) + aByte
	fmt.Println("float64(anInt) + aFloat = ", mixingSum1)
	fmt.Println("anInt + int(aFloat) = ", mixingSum2)
	fmt.Println("anInt + int(aByte) = ", mixingSum3)
	fmt.Println("byte(anInt) + aByte = ", mixingSum4)

	// No truthy values:
	var aFalsyInt int = 0
	var aFalsyString = ""
	// if (!aFalsyInt) {} // Invalid Operation
	if aFalsyInt == 0 && aFalsyString == "" {
		fmt.Println("Falsy values must be compared explicitly")
	}
	var aTruthyInt = 2
	var aTruthyString = "hello"
	// if (aTruthyInt || aTruthyString) {} // Illegal
	if aTruthyInt != 0 && aTruthyString != "" {
		fmt.Println("Truthy values must be compared explicitly")
	}

	// Literals are untyped
	var anIntSum = anInt + 10   // 10 is considerd an int
	var aFloatSum = aFloat + 10 // 10 is considered a float here
	fmt.Println("anInt + 10 (typed as an int) = ", anIntSum)
	fmt.Println("aFloat + 10 (typed as a float) = ", aFloatSum)
	if anInt > 2 {
		fmt.Println("2 is typed as an int for the comparison 'anInt > 2'")
	}
	// var aMixingTypesSum = anInt + "4" // invalid due to type mismatching
	var aMixingTypesSum = anInt + '4'         // rune typed as int
	var anotherMixingtTypesSum = aFloat + '4' // rune typed as float
	fmt.Println("anInt + '4' (the '4' char has value 52) = ", aMixingTypesSum)
	fmt.Println("aFloat + '4' (the '4' char has value 52) = ", anotherMixingtTypesSum)
	var aRune rune = 52 // The correct way would be '4'
	if aRune == '4' {   // This works because
		fmt.Println("aRune (to which we asigned 52) == '4'")
	}

	// var vs :=
	// Var defines a new variable
	var x = 10
	// x := 10 // You can't do this because x is already declared
	x, y := 20, 30 // This reassigns x to 20. Y is created instead.
	fmt.Println("x, y := ", x, ",", y)
	shadowingExample()

	// Constants
	// They are very limited to
	// numeric literals:
	const aNumericConst float64 = 10.2
	// aNumericConst = aNumericConst + 1 // Not compiling: cannot assign to aNumericConst (neither addressable nor a map index expression)
	fmt.Println("aNumericConst = ", aNumericConst)
	// Boolean
	const aBooleanConst bool = true
	fmt.Println("aBooleanConst = ", aBooleanConst)
	// Strings
	const aStringConst string = "Constant"
	fmt.Println("aStringConst = ", aStringConst)
	// Rune
	const aRuneConst rune = 'a'
	fmt.Println("aRuneConst = ", aRuneConst)
	// values returned by complex, real, imag, len, and cap
	const aComplexConst = complex(1.1, 2)
	fmt.Println("aComplexConst = ", aComplexConst)
	// Operations between alreayd defined consts
	const aCalculatedConst = aStringConst + " another string constant"
	fmt.Println("aCalculatedConst = aStringConst + \" another string constant\" -> ", aCalculatedConst)
	// const calculatedConstBasedOnVars = anInt + aTruthyInt // Not compiling:  anInt + aTruthyInt (value of type int) is not constant

	const anUntypedConstant = 10
	var anIntVar int = anUntypedConstant
	var aFloatVar float32 = anUntypedConstant
	var aByteVar byte = anUntypedConstant
	fmt.Println("anUntypedConstant, anIntVar, aFloatVar, aByteVar: ", anUntypedConstant, anIntVar, aFloatVar, aByteVar)

	const aTypedConst float32 = 10.2
	// var anotherIntVar int = aTypedConst // Not compiling: cannot use aTypedConst (constant 10.2 of type float32) as int value in variable declaration

	// unused variables
	var unusedVariable int = 10
	unusedVariable = 20
	fmt.Println("unusedVariable = ", unusedVariable) // reading the variable avoids compiler errors

	// Naming variables
	_1 := 1
	__ := true
	ș := "ș "
	fmt.Println("_1, __, ș: ", _1, __, ș)

	exercise1()
	exercise2()
	exercise3()
}
