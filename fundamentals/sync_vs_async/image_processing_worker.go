package syncvsasync

import (
	"fmt"
	"sync"
	"time"
)

func image_worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(1 * time.Second)
		results <- fmt.Sprintf("Worker %d processed job %d", id, j)
	}
}

func Image_Processing_Worker() {
	fmt.Println("\n===> Image_Processing_Worker")
	jobs := make(chan int, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	// start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go image_worker(w, jobs, results, &wg)
	}

	// send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// wait for workers and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for r := range results {
		fmt.Println(r)
	}
}
