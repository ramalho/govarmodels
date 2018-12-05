package main

import (
	"fmt"
	"unsafe"
)


func InspectString(aString string) {

	fmt.Println("aString:")
	fmt.Printf("\t%#v\n", aString)

	// Get stringPtr of string structure
	stringPtr := unsafe.Pointer(&aString)
	ptrSize := unsafe.Sizeof(stringPtr)

	// Compute address of len
	lenAddr := uintptr(stringPtr) + ptrSize

	// Create pointer to len
	lenPtr := (*int)(unsafe.Pointer(lenAddr))

	fmt.Println("\naString:")

	dataPtr := (*byte)(unsafe.Pointer(*(*uintptr)(stringPtr)))

	fmt.Printf("\t@%p: data %T = %p\n", stringPtr, dataPtr, dataPtr)
	fmt.Printf("\t@%p: len %T = %d\n", lenPtr, *lenPtr, *lenPtr)

	fmt.Println("\ndata:")

	for index := 0; index < len(aString); index++ {
		byteAddr := uintptr(unsafe.Pointer(dataPtr)) + uintptr(index)
		bytePtr := (*byte)(unsafe.Pointer(byteAddr))
		fmt.Printf("\t@%p: %T = %d\n",
			bytePtr, *bytePtr, *bytePtr)
	}


}

func main() {
	aString := "ABC"

	InspectString(aString)
}