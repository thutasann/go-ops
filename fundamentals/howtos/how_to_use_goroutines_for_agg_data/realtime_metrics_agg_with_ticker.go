package howtousegoroutinesforaggdata

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ticker_worker(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		val := rand.Intn(100)
		ch <- val
	}
}

func RealTime_Metrics_Agg_With_Ticker() {
	fmt.Println("\n===> Real Time Metrics Aggregator with Ticker")
	rand.Seed(time.Now().UnixNano())
	metrics := make(chan int, 100)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go ticker_worker(i, metrics, &wg)
	}

	// Aggregator goroutine
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		batch := []int{}
		for {
			select {
			case val, ok := <-metrics:
				if !ok {
					return
				}
				batch = append(batch, val)
			case <-ticker.C:
				if len(batch) > 0 {
					sum := 0
					for _, val := range batch {
						sum += val
					}
					fmt.Printf("📊 Metrics batch: %v, avg=%d\n", batch, sum/len(batch))
					batch = []int{} // reset
				}
			}
		}
	}()

	wg.Wait()
	close(metrics)
	time.Sleep(2 * time.Second)
}
