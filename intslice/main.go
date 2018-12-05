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
	arrayPtr := (*[5]int)(unsafe.Pointer(*(*uintptr)(slicePtr)))

	fmt.Println("intSlice:")

	fmt.Printf("\t@%p: data %T = %p\n", slicePtr, arrayPtr, arrayPtr)

	fmt.Printf("\t@%p: len %T = %d\n", lenPtr, *lenPtr, *lenPtr)

	fmt.Printf("\t@%p: cap %T = %d\n", capPtr, *capPtr, *capPtr)

	fmt.Println("data:")

	fmt.Printf("\t@%p: %T\n", arrayPtr, *arrayPtr)

	for index := 0; index < *capPtr; index++ {
		fmt.Printf("\t\t@%p: [%d] %T = %d\n",
			&(*arrayPtr)[index], index, (*arrayPtr)[index], (*arrayPtr)[index])
	}

}

func main() {
	intSlice := make([]int, 3, 5)
	intSlice[0] = 11
	intSlice[1] = 12
	intSlice[2] = 13

	InspectSlice(intSlice)
}