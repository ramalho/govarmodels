package main

import (
	"fmt"
	"unsafe"
)

func InspectSlice(intSlice []int) {

	fmt.Println("intSlice:")
	fmt.Printf("\t%#v\n\n", intSlice)

	// Get slicePtr of slice structure
	slicePtr := unsafe.Pointer(&intSlice)
	ptrSize := unsafe.Sizeof(slicePtr)

	// Compute addresses of len and cap
	lenAddr := uintptr(slicePtr) + ptrSize
	capAddr := uintptr(slicePtr) + (ptrSize * 2)

	// Create pointers to len and cap
	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	// Get pointer to underlying array
	// How to do this without hardcoding the array size?
	arrayPtr := (*[100]int)(unsafe.Pointer(*(*uintptr)(slicePtr)))

	fmt.Println("intSlice:")

	// Not using %T on next line to show expected data array size
	// fmt.Printf("\t@%p: data %T = %p\n", slicePtr, arrayPtr, arrayPtr)
	fmt.Printf("\t@%p: data *[%d]int = %p\n", slicePtr, *capPtr, arrayPtr)

	fmt.Printf("\t@%p: len %T = %d\n", lenPtr, *lenPtr, *lenPtr)

	fmt.Printf("\t@%p: cap %T = %d\n", capPtr, *capPtr, *capPtr)

	fmt.Println("data:")

	for index := 0; index < *capPtr; index++ {
		fmt.Printf("\t@%p: [%d] %T = %d\n",
			&(*arrayPtr)[index], index, (*arrayPtr)[index], (*arrayPtr)[index])
	}

}

func main() {
	intSlice := make([]int, 3, 5)
	intSlice[0] = 11
	intSlice[1] = 12
	intSlice[2] = 13

	InspectSlice(intSlice)

	for i := 1; i < 5; i++ {
		intSlice = append(intSlice, 10*(13+i))
		InspectSlice(intSlice)
	}

}
