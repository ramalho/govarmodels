package main

import (
	"fmt"
	"unsafe"
)


func InspectString(aString string) {

	fmt.Println("aString:")
	fmt.Printf("\t%#v\n\n", aString)

	// Get stringPtr of slice structure
	stringPtr := unsafe.Pointer(&aString)
	ptrSize := unsafe.Sizeof(stringPtr)

	// Get address of len and cap
	lenAddr := uintptr(stringPtr) + ptrSize

	// Create pointers to len and cap
	lenPtr := (*int)(unsafe.Pointer(lenAddr))

	fmt.Println("aString:")

	// how to express the type of the `data` byte array without hardcoding the len (3) here?
	dataPtr := (*[3]byte)(unsafe.Pointer(*(*uintptr)(stringPtr)))

	fmt.Printf("\t%p→ data %T = %p→\n", stringPtr, dataPtr, dataPtr)
	fmt.Printf("\t%p→ len %T = %d\n", lenPtr, *lenPtr, *lenPtr)

	fmt.Println()

	fmt.Printf("\t%p→ %T\n", dataPtr, *dataPtr)

	for index := 0; index < len(aString); index++ {
		fmt.Printf("\t\t[%d] %p→ %d\n",
			index, &(*dataPtr)[index], (*dataPtr)[index])
	}


}

func main() {
	aString := "ABC"

	InspectString(aString)
}