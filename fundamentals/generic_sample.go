package main

import "fmt"

func Generic_Sample() {
	var t1 int = 123
	fmt.Printf("plusOne: %d\n", plusOne(t1))

	var t2 float64 = 123.12
	fmt.Printf("plusOne: %v\n", plusOne(t2))
}

func plusOne[V int | float64 | int64 | float32 | int32](t V) V {
	return t + 1
}
