package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var anInt int
	var aString string
	var anArray [3]float32
	var anArrayPtr *[3]float32
	var aSlice []float32
	var aMap map[string]int
	var aChannel chan byte

	fmt.Printf("%T (%d): %#v\n", aString, unsafe.Sizeof(aString), aString)
	fmt.Printf("%T (%d): %#v\n", anInt, unsafe.Sizeof(anInt), anInt)
	fmt.Printf("%T (%d): %#v\n", anArray, unsafe.Sizeof(anArray), anArray)
	fmt.Printf("%T (%d): %#v\n", anArrayPtr, unsafe.Sizeof(anArrayPtr), anArrayPtr)
	fmt.Printf("%T (%d): %#v\n", aSlice, unsafe.Sizeof(aSlice), aSlice)
	fmt.Printf("%T (%d): %#v\n", aMap, unsafe.Sizeof(aMap), aMap)
	fmt.Printf("%T (%d): %#v\n", aChannel, unsafe.Sizeof(aChannel), aChannel)


}