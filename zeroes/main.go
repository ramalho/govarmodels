package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var anInt int
	var aFloat32 float32
	var aString string
	var anArray [3]float32
	var anArrayPtr *[3]float32
	var aSlice []float32
	var aMap map[string]int
	var aChannel chan byte

	fmt.Printf("%T\t%d\t%#v\n", anInt, unsafe.Sizeof(anInt), anInt)
	fmt.Printf("%T\t%d\t%#v\n", aFloat32, unsafe.Sizeof(aFloat32), aFloat32)
	fmt.Printf("%T\t%d\t%#v\n", anArray, unsafe.Sizeof(anArray), anArray)
	fmt.Printf("%T\t%d\t%#v\n", aString, unsafe.Sizeof(aString), aString)
	fmt.Printf("%T\t%d\t%#v\n", anArrayPtr, unsafe.Sizeof(anArrayPtr), anArrayPtr)
	fmt.Printf("%T\t%d\t%#v\n", aSlice, unsafe.Sizeof(aSlice), aSlice)
	fmt.Printf("%T\t%d\t%#v\n", aMap, unsafe.Sizeof(aMap), aMap)
	fmt.Printf("%T\t%d\t%#v\n", aChannel, unsafe.Sizeof(aChannel), aChannel)

}
