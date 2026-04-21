package main

import (
	"fmt"
	"reflect"
	"slices"
)

func printSection(sectionTitle string) {
	fmt.Println()
	fmt.Println("#################################################")
	fmt.Println("#", sectionTitle)
	fmt.Println("#################################################")
	fmt.Println()
}

func customSliceCompare(a int, b string) bool {
	switch a {
	case 1:
		return b == "a"
	case 2:
		return b == "b"
	case 3:
		return b == "cd"
	default:
		return false
	}
}

func example37() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:] // This picks up only "c" and "d" plus an available slot.
	fmt.Println("capacities:", cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x) // x = [a b i j y]
	fmt.Println("y:", y) // y = [a b i j y]
	fmt.Println("z:", z) // z = [i j y]
}

func example38() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d") // length 4, capacity 5
	y := x[:2:2]                      // length 2, capacity 2 - 0 = 2
	z := x[2:4:4]                     // length 2, capacity 4 - 2 = 2
	fmt.Println("capacities:", cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k") // new backing array since it exceeds capacity
	x = append(x, "x")           // same backing array as it has space
	z = append(z, "y")           // new backing array since it exceeds capacity
	fmt.Println("x:", x)         // x = [a b c d x]
	fmt.Println("y:", y)         // y = [a b i j k]
	fmt.Println("z:", z)         // z = [c d y]
}

func main() {
	// Arrays
	printSection("Arrays")

	var arrayVar [3]int
	var arrayLiteralVar = [3]int{1, 2, 3}
	var arraySparseVar = [12]int{1, 5: 4, 6, 10: 100, 15}
	var arraySparseVar2 = [15]int{1, 5: 4, 6, 10: 100, 15}
	var arrayOfDerivedLength = [...]int{1, 2, 3}
	fmt.Println("arrayVar: ", arrayVar)
	fmt.Println("arrayLiteralVar: ", arrayLiteralVar)
	fmt.Println("arraySparseVar: ", arraySparseVar)
	fmt.Println("arraySparseVar2: ", arraySparseVar2)
	fmt.Println("arrayOfDerivedLength: ", arrayOfDerivedLength)

	fmt.Println()

	if arrayLiteralVar == arrayOfDerivedLength {
		fmt.Println("Arrays are equal if they have the same length and the same elements)", arrayLiteralVar, arrayOfDerivedLength)
	}

	var arrayDifferentOrder = [...]int{1, 3, 2}
	var arrayDifferentSize = [...]int{1, 3}

	if arrayLiteralVar != arrayDifferentOrder {
		fmt.Println("Arrays are different if they have the same length and the same elements)", arrayLiteralVar, arrayDifferentOrder, arrayDifferentSize)
	}
	// if arrayLiteralVar != arrayDifferentSize {} // compiler error

	var multidimensionalArray [2][3]int
	fmt.Println()
	fmt.Println("multidimensionalArray :", multidimensionalArray)

	// Reading and writing
	fmt.Println("Reading array at index 0: ", arrayLiteralVar[0])
	arrayLiteralVar[0] = 20
	fmt.Println("New element at index 0: ", arrayLiteralVar[0])

	//  arrayLiteralVar[-1] // Compile timer error: invalid argument: index -1 (constant of type int) must not be negative (vet)
	// var outOfBoundsRead = arrayLiteralVar[len(arrayLiteralVar)+1] // Compile time error (vet)
	// var length = len(arrayLiteralVar)
	// length += 1
	// var outOfBoundsRead = arrayLiteralVar[length] // Causes panic
	// fmt.Println(outOfBoundsRead)

	printSection("Slices")

	var sliceVar = []int{10, 20, 30}
	var sliceSparse = []int{1, 2, 5: 4, 6, 10: 20, 15: 30} // Fills in the indexes not specified with default value, i.e. 0 in this case.
	var multidimensionalSlice [][]int                      // Multidimensional slice.
	var emptySlice []int                                   // initialized to nil
	fmt.Println("sliceVar: ", sliceVar)
	fmt.Println("sliceSparse: ", sliceSparse)
	fmt.Println("multidimensionalSlice: ", multidimensionalSlice)
	fmt.Println("emptySlice: ", emptySlice)
	fmt.Println()

	// if emptySlice == sliceVar {} // vet error: invalid operation: emptySlice == sliceVar (slice can only be compared to nil)

	if emptySlice == nil {
		fmt.Println("Empty slice: (emptySlice == nil)")
	}

	var firstSlice = []int{1, 2, 3}
	var secondSlice = []int{1, 2, 3}
	var thirdSlice = []int{1, 2}
	var stringSlice = []string{"a", "b", "cd"}
	fmt.Println("firstSlice: ", firstSlice)
	fmt.Println("secondSlice: ", secondSlice)
	fmt.Println("thirdSlice: ", thirdSlice)
	fmt.Println("stringSlice: ", stringSlice)
	if slices.Equal(firstSlice, secondSlice) {
		fmt.Println("firstSlice equals secondSlice")
	}
	if !slices.Equal(firstSlice, thirdSlice) {
		fmt.Println("firstSlice is different from thirdSlice")
	}
	// slices.Equal(firstSlice, stringSlice) // vet Error: in call to slices.Equal, type []string of stringSlice does not match inferred type []int for S
	if slices.EqualFunc(firstSlice, stringSlice, customSliceCompare) {
		fmt.Println("slices are equal through custom comparison function")
	}
	fmt.Println()

	fmt.Println("Length of firstSlice:", len(firstSlice))
	fmt.Println("Length of emptySlice", len(emptySlice))
	fmt.Println()

	fmt.Println("thirdSlice before append:", thirdSlice)
	var fourthSlice = append(thirdSlice, 4) // Since Go is call by value a copy of third slice is made and passed. It does not mutate the original value.
	fmt.Println("thirdSlice, fourthSlice:", thirdSlice, fourthSlice)
	var fifthSlice = append(thirdSlice, fourthSlice...) // use the expand operator to add a slice to another one.
	fmt.Println("fifthSlice = append(thirdSlice, fourthSlice...):", fifthSlice)
	fmt.Println()

	// Slice capacity
	var capSl []int
	fmt.Println("capSl info:", capSl, len(capSl), cap(capSl))
	capSl = append(capSl, 10)
	fmt.Println("capSl info:", capSl, len(capSl), cap(capSl))
	capSl = append(capSl, 20)
	fmt.Println("capSl info:", capSl, len(capSl), cap(capSl))
	capSl = append(capSl, 30)
	fmt.Println("capSl info:", capSl, len(capSl), cap(capSl))
	capSl = append(capSl, 40)
	fmt.Println("capSl info:", capSl, len(capSl), cap(capSl))
	capSl = append(capSl, 50)
	fmt.Println("capSl info:", capSl, len(capSl), cap(capSl))
	fmt.Println()

	var newSlice = make([]int, 5)
	fmt.Println("newSlice: ", newSlice, len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 1) // This does not populate the slice. It add a new element.
	fmt.Println("newSlice after append: ", newSlice, len(newSlice), cap(newSlice))
	var largeCapSlice = make([]int, 4, 120)
	fmt.Println("largeCapSlice: ", largeCapSlice, len(largeCapSlice), cap(largeCapSlice))
	var emptyCapSlice = make([]int, 0, 200)
	// Make always returns a slice which is different that nil. len(nil) == len(emptySlice)
	fmt.Println("emptyCapSlice: ", emptyCapSlice, len(emptyCapSlice), cap(emptyCapSlice), emptyCapSlice == nil)
	emptyCapSlice = append(emptyCapSlice, 1, 2, 3, 4)
	fmt.Println("emptyCapSlice after append: ", emptyCapSlice, len(emptyCapSlice), cap(emptyCapSlice), emptyCapSlice == nil)
	clear(emptyCapSlice)
	fmt.Println("emptyCapSlice after clear: ", emptyCapSlice, len(emptyCapSlice), cap(emptyCapSlice), emptyCapSlice == nil)
	var strSlice = []string{"a", "b", "c"}
	fmt.Println("strSlice: ", strSlice, len(strSlice), cap(strSlice))
	clear(strSlice)
	fmt.Println("strSlice after clear: ", strSlice, len(strSlice), cap(strSlice))
	fmt.Println()

	// Declaring slices
	var data []int                                                                                                // nil Slice: prefer this for slices that might not contain any data, no memory allocation
	var noData = []int{}                                                                                          // Empty Slice. Different from nil slice
	fmt.Println("data equals noData:", data, noData, slices.Equal(data, noData), reflect.DeepEqual(data, noData)) // slices.Equal returns true nil and empty slice are functionally similar, reflec.DeepEqual distinguishes them
	var someData = []int{1, 2, 3}                                                                                 // Use literals when you have some values you're aware of when writing the code.
	fmt.Println("var someData = []int{1, 2, 3} -> Known elements from the beggining: ", someData)                 //
	var bufferSlice = make([]int, 30)                                                                             // If you're using it as a buffer, specify the length.
	fmt.Println("var bufferSlice = make([]int, 30) -> for buffers", bufferSlice)                                  //
	var knownSize = make([]int, 3)                                                                                // For known size initialize the size with make
	knownSize[0] = 1                                                                                              //
	knownSize[1] = 2                                                                                              //
	fmt.Println("known size -> var knownSize = make([]int, 3): ", knownSize)                                      // Downside: it has appended zero if all elements are not filled in
	// knownSize[3] = 3                                                                                           // Downside: this will cause a panic
	var knownSizeAlt = make([]int, 0, 30)                                               // Use append to add elements
	knownSizeAlt = append(knownSizeAlt, 1, 2, 3)                                        // No appended 0s or panics due to assignment
	fmt.Println("known size alternative, specify capacity with 0 length", knownSizeAlt) // Downside: slower but less likely to introduce bugs.
	fmt.Println()

	// Slicing slices

	var mainSlice = []string{"a", "b", "c", "d"}
	var firstSubslice = mainSlice[:2]
	var secondSubslice = mainSlice[1:]
	var thirdSubslice = mainSlice[1:3]
	var fourthSubslice = mainSlice[:]
	fmt.Println("mainSlice: ", mainSlice)
	fmt.Println("firstSubslice: ", firstSubslice)
	fmt.Println("secondSubslice: ", secondSubslice)
	fmt.Println("thirdSubslice: ", thirdSubslice)
	fmt.Println("fourthSubslice: ", fourthSubslice)
	fmt.Println()

	fmt.Println("modifying mainSlice at index 1")
	mainSlice[1] = "x"
	fmt.Println("mainSlice:", mainSlice)
	fmt.Println("firstSubslice: ", firstSubslice)
	fmt.Println("secondSubslice: ", secondSubslice)
	fmt.Println()

	fmt.Println("modifying firstSubslice at index 0")
	firstSubslice[0] = "y"
	fmt.Println("mainSlice:", mainSlice)
	fmt.Println("firstSubslice: ", firstSubslice)
	fmt.Println("secondSubslice: ", secondSubslice)
	fmt.Println()

	fmt.Println("modifying secondSubslice at index 1")
	secondSubslice[1] = "z"
	fmt.Println("mainSlice:", mainSlice)
	fmt.Println("firstSubslice: ", firstSubslice)
	fmt.Println("secondSubslice: ", secondSubslice)
	fmt.Println()

	fmt.Println("Appending to a firstSubslice")
	firstSubslice = append(firstSubslice, "w") // this has length 2 and capacity 2. When appending it extends to the capacity of the mainSlice.
	fmt.Println("mainSlice:", mainSlice)
	fmt.Println("firstSubslice: ", firstSubslice)
	fmt.Println("secondSubslice: ", secondSubslice)
	fmt.Println()

	fmt.Println("Appending to a secondSubslice")
	secondSubslice = append(secondSubslice, "p") // this has length 3 and capacity 3. When appending to it, a new array gets created. It diverges from the main slice array.
	fmt.Println("mainSlice:", mainSlice)
	fmt.Println("firstSubslice: ", firstSubslice)
	fmt.Println("secondSubslice: ", secondSubslice)
	fmt.Println()

	fmt.Println("Appending to the mainSlice")
	mainSlice = append(mainSlice, "o") // this has len 4 and cap 4 so its capacity gets increased and a new array gets generated. This will make the mainSlice diverge from the firstSubslice.
	fmt.Println("mainSlice:", mainSlice)
	fmt.Println("firstSubslice: ", firstSubslice)
	fmt.Println("secondSubslice: ", secondSubslice)
	fmt.Println()

	fmt.Println("Modifying the first and secondSubslice")
	secondSubslice[0] = "t" // Independent from the mainSlice
	firstSubslice[0] = "l"  // Independent from the mainSlice
	fmt.Println("mainSlice:", mainSlice)
	fmt.Println("firstSubslice: ", firstSubslice)
	fmt.Println("secondSubslice: ", secondSubslice)
	fmt.Println()

	example37() // In the examples before the capacity is not sufficient so new arrays are created. In the example the capacity is sufficient, thus no array is being created so data is overwritten.
	fmt.Println()

	example38()
	fmt.Println()

	// copy
	var srcSlice = []int{1, 2, 3, 4}
	var destSlice = make([]int, 4)
	fmt.Println("before copy: srcSlice, destSlice", srcSlice, destSlice)
	var num = copy(destSlice, srcSlice) // Number of elements that are copied. Depends on the minimum of the length of the two slices. Capacity is not relevant here.
	fmt.Println("srcSlice, destSlice, num", srcSlice, destSlice, num)
	fmt.Println()

	destSlice = make([]int, 2)
	fmt.Println("before copy: srcSlice, destSlice", srcSlice, destSlice)
	num = copy(destSlice, srcSlice)
	fmt.Println("srcSlice, destSlice, num", srcSlice, destSlice, num)
	fmt.Println()

	// Copy from a part of the slice
	destSlice = make([]int, 2)
	fmt.Println("before copy: srcSlice, destSlice", srcSlice, destSlice)
	num = copy(destSlice, srcSlice[2:])
	fmt.Println("srcSlice, destSlice, num", srcSlice, destSlice, num)
	fmt.Println()

	// Overwrite slices of the same slice
	fmt.Println("before copy: srcSlice", srcSlice)
	num = copy(srcSlice[:3], srcSlice[1:])
	fmt.Println("srcSlice, num", srcSlice, num)
	fmt.Println()

	srcSlice = []int{1, 2, 3, 4}
	var destArray = [4]int{5, 6, 7, 8}
	var destSlice2 = make([]int, 2)
	copy(destSlice2, destArray[:])
	fmt.Println("destSlice2:", destSlice2)
	copy(destArray[:], srcSlice)
	fmt.Println("destArray:", destArray)
	fmt.Println()

	// Converting arrays to slices
	var anArray = [4]int{1, 2, 3, 4}
	var sliceFromArray = anArray[:]
	var subsliceFromArray = anArray[:2]
	sliceFromArray[0] = 10    // same memory sharing
	subsliceFromArray[1] = 20 // same memory sharing
	fmt.Println("anArray, sliceFromArray, subsliceFromArray", anArray, sliceFromArray, subsliceFromArray)
	fmt.Println()

	var sliceToConvert = []int{1, 2, 3, 4}
	var arrFromSlice = [4]int(sliceToConvert)
	var smallArrFromSlice = [2]int(sliceToConvert)
	// var panicArrFromSlice = [5]int(sliceToConvert) // Panic panic: runtime error: cannot convert slice with length 4 to array or pointer to array with length 5
	sliceToConvert[0] = 10
	smallArrFromSlice[0] = 20
	fmt.Println("sliceToConvert, arrFromSlice, smallArrFromSlice:", sliceToConvert, arrFromSlice, smallArrFromSlice)
	fmt.Println()

	// Pointer conversion
	var initSlice = []int{1, 2, 3, 4}
	var arrayPointer = (*[4]int)(initSlice)
	initSlice[0] = 10
	arrayPointer[1] = 20
	fmt.Println("initSlice, arrayPointer", initSlice, arrayPointer)
	fmt.Println()

	// Strings, runes and bytes
	var str string = "Hello world"
	var b byte = str[6]
	fmt.Println("str, b:", str, b)
	var bSlice1 = str[4:7]
	var bSlice2 = str[:5]
	var bSlice3 = str[6:]
	fmt.Println("bSlice1, bSlice2, bSlice3:", bSlice1, bSlice2, bSlice3)
	fmt.Println()

	// Mulitple bytes encoding
	var multiByteStr = "Hello 🌞"
	var bSlice4 = multiByteStr[4:7]
	var bSlice5 = multiByteStr[:5]
	var bSlice6 = multiByteStr[6:]
	fmt.Println("multiByteStr, len:", multiByteStr, len(multiByteStr))
	fmt.Println("bSlice4:", bSlice4)
	fmt.Println("bSlice5:", bSlice5)
	fmt.Println("bSlice6:", bSlice6)

	var byteSlice []byte = []byte(multiByteStr)
	var runeSlice []rune = []rune(multiByteStr)
	fmt.Println("byteSlice:", byteSlice)
	fmt.Println("runeSlice:", runeSlice) // uncommon

	// Maps
}
