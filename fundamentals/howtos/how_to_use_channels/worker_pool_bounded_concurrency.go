package howtousechannels

import (
	"fmt"
	"sync"
	"time"
)

func bounded_concurrency_worker(id int, jobs <-chan int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	for j := range jobs {
		// acquire token
		sem <- struct{}{}
		fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("worker %d finished job %d\n", id, j)
		<-sem // released token
	}
}

// sem bounds concurrency to 3 simultaneous active jobs, regardless of number of worker goroutines.
func Bounded_Concurrency() {
	fmt.Println("\n===> Worker Pool (Bounded_Concurrency) -- use buffered semaphore")
	jobCnt := 10
	jobs := make(chan int, jobCnt)
	sem := make(chan struct{}, 3) // allow 3 concurrently running tasks
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go bounded_concurrency_worker(i, jobs, &wg, sem)
	}

	for j := 0; j < jobCnt; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}
