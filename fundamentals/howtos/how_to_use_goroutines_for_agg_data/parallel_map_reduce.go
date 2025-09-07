package howtousegoroutinesforaggdata

import (
	"fmt"
	"sync"
)

func parallel_map_reduce_worker(id int, nums []int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Printf("Worker %d processed: %v => sum=%d\n", id, nums, sum)
	out <- sum
}

func Parallel_Map_Reduce_Sample() {
	fmt.Println("\n===> Parallel Map Reduce Sample")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	chunkSize := 3
	var wg sync.WaitGroup
	out := make(chan int, len(data)/chunkSize)

	// split work into chunks
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go parallel_map_reduce_worker(i/chunkSize, data[i:end], out, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	total := 0
	for sum := range out {
		total += sum
	}
	fmt.Println("Final aggregated sum:", total)
}
