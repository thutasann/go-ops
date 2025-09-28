package howtoslices

import "fmt"

func slidingWindow[T any](s []T, size int) [][]T {
	var result [][]T
	for i := 0; i+size <= len(s); i++ {
		result = append(result, s[i:i+size])
	}

	return result
}

func Sliding_Window_Sample() {
	fmt.Println("\n==> Sliding Window Sample")
	nums := []int{1, 2, 3, 4, 5}
	// [[1 2 3] [2 3 4] [3 4 5]]
	fmt.Println(slidingWindow(nums, 3))
}
