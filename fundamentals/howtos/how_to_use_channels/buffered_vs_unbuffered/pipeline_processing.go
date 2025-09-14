package bufferedvsunbuffered

import (
	"fmt"
	"sync"
	"time"
)

func pipeline_producer(out chan<- int) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func pipeline_worker(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		fmt.Printf("Worker %d processing %d\n", id, num)
		time.Sleep(500 * time.Millisecond)
		out <- num * num
	}
}

func pipeline_collector(in <-chan int, done chan<- struct{}) {
	for res := range in {
		fmt.Printf("Collected result: %d\n", res)
	}
	done <- struct{}{}
}

func Pipeline_Processing() {
	fmt.Println("\n==> Pipeline Processing")

	jobs := make(chan int)
	results := make(chan int, 10)
	done := make(chan struct{})

	var wg sync.WaitGroup

	// producer
	go pipeline_producer(jobs)

	// workers (fan-out)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go pipeline_worker(w, jobs, results, &wg)
	}

	// Collector (fan-in)
	go func() {
		wg.Wait()
		close(results)
	}()

	go pipeline_collector(results, done)

	<-done
	fmt.Println("🚀 Pipeline finished")
}
