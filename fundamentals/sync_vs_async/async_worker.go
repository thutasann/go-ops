package syncvsasync

import (
	"fmt"
)

func async_worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * 2
	}
}

func Async_Worker_Sample() {
	fmt.Println("\n==> Async Worker Sample")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go async_worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= 5; r++ {
		fmt.Println("Result:", <-results)
	}
}
