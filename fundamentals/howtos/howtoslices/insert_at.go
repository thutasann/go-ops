package howtoslices

import "fmt"

func insert_at[T any](s []T, index int, value T) []T {
	if index < 0 || index > len(s) {
		return s
	}
	s = append(s[:index], append([]T{value}, s[index:]...)...)
	return s
}

func Insert_At_Sample() {
	fmt.Println("\n===> Insert At Sample")
	nums := []int{1, 2, 3, 4}
	nums = insert_at(nums, 2, 99)
	fmt.Println(nums) // [1 2 99 3 4]
}
