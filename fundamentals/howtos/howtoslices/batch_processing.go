package howtoslices

import "fmt"

func Batch_Processing[T any](data []T, size int) [][]T {
	var batches [][]T
	for size < len(data) {
		data, batches = data[size:], append(batches, data[:size:size])
	}
	batches = append(batches, data)

	return batches
}

func Batch_Processing_Sample() {
	fmt.Println("\n==> Batch Processing Sample")

	data := make([]int, 23)
	for i := range data {
		data[i] = i + 1
	}
	for i, batch := range Batch_Processing(data, 10) {
		fmt.Printf("Batch %d: %v\n", i+1, batch)
	}
}
