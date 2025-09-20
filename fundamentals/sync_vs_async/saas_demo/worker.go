package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(1 * time.Second)
		results <- fmt.Sprintf("Worker %d processed job %d", id, j)
	}
}

func startWorkerPool() {
	jobs := make(chan int, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	// start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Feed jobs async
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Collect results async
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	go func() {
		for r := range results {
			fmt.Println("[Worker Pool]:", r)
		}
	}()
}
