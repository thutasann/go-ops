package howtousegoroutinesforaggdata

import (
	"fmt"
	"sync"
)

func Sum_Of_Numbers() {
	fmt.Println("\n===> Sum Of Numbers")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	chunkSize := 2
	var wg sync.WaitGroup
	results := make(chan int, len(nums)/chunkSize)

	// split work into chunks
	for i := 0; i < len(nums); i += chunkSize {
		end := i + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()
			sum := 0
			for _, n := range chunk {
				sum += n
			}
			results <- sum
		}(nums[i:end])
	}

	// Close results when done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Aggregate results
	total := 0
	for r := range results {
		total += r
	}
	fmt.Println("Total sum:", total)
}
