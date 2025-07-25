package main

import "fmt"

func Generic_Sample() {
	var t1 int = 123
	fmt.Printf("plusOne: %d\n", plusOne(t1))

	var t2 float64 = 123.12
	fmt.Printf("plusOne: %v\n", plusOne(t2))
}

func Generic_Sample_Two() {
	ints := []int{1, 2, 3, 4}
	evens := filter(ints, func(i int) bool { return i%2 == 0 })
	fmt.Println(evens) // [2 4]
}

func plusOne[V int | float64 | int64 | float32 | int32](t V) V {
	return t + 1
}

func filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}
