package howtoslices

import "fmt"

func rotate_left[T any](s []T, k int) []T {
	n := len(s)
	k = k % n
	return append(s[k:], s[:k]...)
}

// Rotae Slice Sample
func Rotate_Slice_Sample() {
	fmt.Println("\n===> Rotate Slice Sample")
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(rotate_left(data, 2))
}
