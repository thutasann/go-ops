package howtoslices

import (
	"fmt"
)

func chunk_data[T any](s []T, size int) [][]T {
	var chunks [][]T
	for size < len(s) {
		s, chunks = s[size:], append(chunks, s[:size:size])
	}
	chunks = append(chunks, s)
	return chunks
}

func Chunking_Data_Sample() {
	fmt.Println("\n===> Chunking Data Sample")
	data := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(chunk_data(data, 3)) // [[1 2 3] [4 5 6] [7]]
}
