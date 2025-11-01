package advancedconcepts

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Out   string
	Err   error
}

func context_worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return // context cancelled -> exit
		case j, ok := <-jobs:
			if !ok {
				return // jobs channel closed
			}
			d := time.Duration(rand.Intn(300)+100) * time.Millisecond
			select {
			case <-ctx.Done():
				return
			case <-time.After(d):
			}

			// occasionally simulate an error
			if rand.Intn(20) == 0 {
				results <- Result{JobID: j.ID, Err: errors.New("transient failure")}
			} else {
				results <- Result{JobID: j.ID, Out: fmt.Sprintf("worker %d processed job %d in %v", id, j.ID, d)}
			}
		}
	}
}

func WorkerPool_With_Context_Cancellation() {
	rand.Seed(time.Now().UnixNano())
	const numWorkers = 5
	const numJobs = 50

	jobs := make(chan Job)
	results := make(chan Result)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go context_worker(ctx, w, jobs, results, &wg)
	}

	// producer goroutine
	go func() {
		for i := 1; i <= numJobs; i++ {
			select {
			case <-ctx.Done():
				break
			case jobs <- Job{ID: i}:
			}
		}
		close(jobs)
	}()

	// aggregator: read results until workers done or context cancelled
	go func() {
		wg.Wait()
		close(results)
	}()

	// consume
	success := 0
	fail := 0
	for r := range results {
		if r.Err != nil {
			fail++
			fmt.Printf("[ERR] job %d: %v\n", r.JobID, r.Err)
			// Example: cancel everything on first error (optional policy)
			// cancel()
		} else {
			success++
			fmt.Println(r.Out)
		}
	}

	fmt.Printf("done: success=%d fail=%d\n", success, fail)

}
