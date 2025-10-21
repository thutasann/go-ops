package icselect

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func pool_worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("👷 Worker %d processing job %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
		results <- job * 2
	}
}

func Worker_Pool_Select() {
	fmt.Println("\n==> Worker Pool Select")

	rand.Seed(time.Now().UnixNano())
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	numWorkers := 3
	var wg sync.WaitGroup

	// start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go pool_worker(i, jobs, results, &wg)
	}

	// send jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for res := range results {
		fmt.Printf("✅ Result: %d\n", res)
	}
}
