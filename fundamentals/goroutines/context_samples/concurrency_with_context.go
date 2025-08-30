package contextsamples

import (
	"context"
	"fmt"
	"time"
)

func context_worker(ctx context.Context, jobs <-chan int, results chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker cancelled")
			return
		case j := <-jobs:
			time.Sleep(time.Millisecond * 500) // simulate work
			results <- j * 2
		}
	}
}

func Concurrency_With_Context_Sample() {
	fmt.Println("\n===> Concurrency_With_Context_Sample")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs := make(chan int)
	results := make(chan int)

	go context_worker(ctx, jobs, results)

	for i := 0; i < 5; i++ {
		jobs <- i
	}

	for i := 0; i < 5; i++ {
		select {
		case res := <-results:
			fmt.Println("Result:", res)
		case <-ctx.Done():
			fmt.Println("Main timeout reached")
			return
		}
	}
}
