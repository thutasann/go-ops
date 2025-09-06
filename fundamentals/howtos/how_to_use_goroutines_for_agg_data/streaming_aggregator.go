package howtousegoroutinesforaggdata

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Metric struct {
	Service string
	Value   int
}

func Streaming_Aggregator() {
	fmt.Println("\n==> Streaming Aggregator")
	rand.Seed(time.Now().UnixNano())

	metricsChan := make(chan Metric, 100)
	aggregated := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Simulate services pushing metrics
	services := []string{"auth", "payment", "orders"}
	for _, s := range services {
		wg.Add(1)
		go func(service string) {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				metricsChan <- Metric{Service: service, Value: rand.Intn(100)}
				time.Sleep(time.Millisecond * 200)
			}
		}(s)
	}

	// Aggregator
	go func() {
		for m := range metricsChan {
			mu.Lock()
			aggregated[m.Service] += m.Value
			mu.Unlock()
		}
	}()

	wg.Wait()
	close(metricsChan)

	time.Sleep(time.Second)
	fmt.Println("Aggregated metrics: ", aggregated)
}
