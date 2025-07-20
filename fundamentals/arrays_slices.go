package main

import "fmt"

// - Arrays hav a fixed length whereas slices are dynamic
// - Arrays are the building block of slices
//
// Array:
// - var buffer [7]byte
// - var arr1 [7]int = [7]int{1,2,3,4,5,6,7}
//
// Slice:
// - var buffer []byte
// - var arr2 []byte = arr1[1:3]
func Array_Slice_Sample() {
	var arr1 [7]int = [7]int{7, 3, 6, 0, 4, 9, 10}
	fmt.Println("arr1 -> ", arr1)
	fmt.Printf("length: %d and cap: %d\n", len(arr1), cap(arr1))

	var arr2 []int = arr1[1:3]
	fmt.Println("arr2 -> ", arr2)
	fmt.Printf("length: %d and cap: %d\n", len(arr2), cap(arr2))

	arr2 = arr2[0 : len(arr2)+2]
	fmt.Println("arr2 after -->", arr2)
	fmt.Printf("length: %d and cap: %d\n", len(arr2), cap(arr2))

	for k := range arr2 {
		arr2[k] += 1
	}

	fmt.Println("arr2 after after -->", arr2)
	fmt.Printf("length: %d and cap: %d\n", len(arr2), cap(arr2))
	fmt.Println("arr 1 after after --> ", arr1)

	var arr3 []int = []int{1, 2, 3}
	fmt.Println("arr3 --> ", arr3)
	fmt.Printf("length: %d and cap: %d\n", len(arr3), cap(arr3))

	arr3 = append(arr3, 5)
	fmt.Println("after append, arr3 --> ", arr3)
	fmt.Printf("length: %d and cap: %d\n", len(arr3), cap(arr3))

	var arr4 = make([]int, 3, 9)
	fmt.Println("arr4 --> ", arr4)
	fmt.Printf("length: %d and cap: %d\n", len(arr4), cap(arr4))
}
