package howtoslices

import (
	"fmt"
	"sync"
)

func slice_worker(id int, jobs []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func Split_Work_For_Goroutines() {
	fmt.Println("\n==> Split Work for Goroutines")
	jobs := []int{1, 2, 3, 4, 5, 6}
	chunkSize := 2
	var wg sync.WaitGroup

	for i := 0; i < len(jobs); i += chunkSize {
		end := i + chunkSize
		if end > len(jobs) {
			end = len(jobs)
		}
		wg.Add(1)
		go slice_worker(1/chunkSize+1, jobs[i:end], &wg)
	}

	wg.Wait()
}
