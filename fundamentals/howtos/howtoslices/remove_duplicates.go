package howtoslices

import "fmt"

func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	results := []int{}

	for _, n := range nums {
		if !seen[n] {
			seen[n] = true
			results = append(results, n)
		}
	}

	return results
}

func Remove_Duplicates() {
	fmt.Println("\n==> Remove Duplicates")
	data := []int{1, 2, 2, 3, 4, 4, 5}
	clean := removeDuplicates(data)
	fmt.Println("Cleaned:", clean) // [1 2 3 4 5]
}
