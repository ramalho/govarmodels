package main

import (
	"fmt"
	"unsafe"
)

func InspectArray(intArray [5]int) {

	fmt.Println("intArray:")
	fmt.Printf("\t%#v\n\n", intArray)

	// Get slicePtr of slice structure
	arrayPtr := unsafe.Pointer(&intArray)

	fmt.Println("intArray:")

	fmt.Printf("\t@%p: %T\n", arrayPtr, intArray)

	for index := 0; index < len(intArray); index++ {
		fmt.Printf("\t\t@%p: [%d] %T = %d\n",
			&(intArray)[index], index, intArray[index], intArray[index])
	}

}


func main() {
	intArray := [5]int{11, 12, 13}

	InspectArray(intArray)
}
