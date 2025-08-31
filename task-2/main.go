package main

import (
	"fmt"
	"math/rand"
	"unsafe"
)

func main() {
	const sliceMsg = "slice is "
	const lengthMsg = "length is "
	slice1 := make([]int, 5, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(3000)
	}
	fmt.Println(sliceMsg, slice1, lengthMsg, len(slice1), cap(slice1))
	slicepointer := (*[3]int)(unsafe.Pointer(&slice1))
	slicepointer[1] = 10
	fmt.Println(sliceMsg, slice1, lengthMsg, len(slice1), cap(slice1))
	for _, v := range slice1 {
		fmt.Println("the data :", v)
	}
	arr1 := [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	slicepointer[0] = int(uintptr(unsafe.Pointer(&arr1[0])))
	fmt.Println(sliceMsg, slice1, lengthMsg, len(slice1), cap(slice1))

}
