package main

import (
	"fmt"
	"unsafe"
)

// Based on Bill Kennedy's InspectSlice:
// https://www.ardanlabs.com/blog/2013/12/three-index-slices-in-go-12.html
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
	arrayPtr := (*[100]int)(unsafe.Pointer(*(*uintptr)(slicePtr))) // ➊

	fmt.Println("intSlice:")

	// %T displays *[100]int because it's the array size hardcoded at ➊
	// fmt.Printf("\t@%p: data %T = %p\n", slicePtr, arrayPtr, arrayPtr)
	// instead, I am using *[%d]int to display the size of the underlying
	// array as given by *capPtr
	fmt.Printf("\t@%p: data *[%d]int = %p\n", slicePtr, *capPtr, arrayPtr)

	fmt.Printf("\t@%p: len %T = %d\n", lenPtr, *lenPtr, *lenPtr)

	fmt.Printf("\t@%p: cap %T = %d\n", capPtr, *capPtr, *capPtr)

	/*
		fmt.Println("data:")

		for index := 0; index < *capPtr; index++ {
			fmt.Printf("\t@%p: [%d] %T = %d\n",
				&(*arrayPtr)[index], index, (*arrayPtr)[index], (*arrayPtr)[index])
		}
	*/

}

func main() {
	intSlice := make([]int, 3, 5)
	intSlice[0] = 11
	intSlice[1] = 12
	intSlice[2] = 13

	InspectSlice(intSlice)

	for _, n := range []int{140, 150, 160} {
		intSlice = append(intSlice, n)
		InspectSlice(intSlice)
	}

}
