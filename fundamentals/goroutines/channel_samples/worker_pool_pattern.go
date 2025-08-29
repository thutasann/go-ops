package channelsamples

import (
	"fmt"
	"sync"
	"time"
)

func worker_pool(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(1 * time.Second)
		results <- job * 2
	}
}

func Worker_Pool_Pattern() {
	fmt.Println("\n====> Worker Pool Pattern")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker_pool(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for r := range results {
		fmt.Println("Result:", r)
	}
}
