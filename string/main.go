package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func InspectString(source, aString string) {

	fmt.Println(source)

	// Make pointer to string structure
	stringPtr := unsafe.Pointer(&aString)
	ptrSize := unsafe.Sizeof(stringPtr)

	// Make pointer to content
	contentPtr := (*byte)(unsafe.Pointer(*(*uintptr)(stringPtr)))

	// Compute address of len field
	lenAddr := uintptr(stringPtr) + ptrSize

	// Make pointer to len field
	lenPtr := (*int)(unsafe.Pointer(lenAddr))

	fmt.Println("\nheader struct:")

	fmt.Printf("\t@%p: data %T = %p\n", stringPtr, contentPtr, contentPtr)
	fmt.Printf("\t@%p: len %T = %d\n", lenPtr, *lenPtr, *lenPtr)

	fmt.Println("\ndata:")

	for index := 0; index < len(aString); index++ {
		// Compute address to one byte of content
		byteAddr := uintptr(unsafe.Pointer(contentPtr)) + uintptr(index)
		// Make pointer to that byte of content
		bytePtr := (*byte)(unsafe.Pointer(byteAddr))
		fmt.Printf("\t@%p: %T = %d\n",
			bytePtr, *bytePtr, *bytePtr)
	}
	fmt.Println(strings.Repeat("─", 60))
}

func main() {
	aString := "ABC"
	aStringCopy := aString
	anotherString := "ABC"

	// Same data
	InspectString(`aString := "ABC"`, aString)
	InspectString(`aStringCopy := aString`, aStringCopy)
	InspectString(`anotherString := "ABC"`, anotherString)

	// Changing data
	InspectString(`aString := "ABC"`, aString)
	aString += "?"
	InspectString(`aString += "?"`, aString)
	aString += "!"
	InspectString(`aString += "!"`, aString)
}
