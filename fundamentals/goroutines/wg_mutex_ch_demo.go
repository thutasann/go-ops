package goroutines

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Amount float64
}

// Imagine you're building a service that processes orders concurrently:
//
// - Orders come in via a channel.
//
// - Workers pull from the channel and process them.
//
// - Total revenue is calculated safely using a mutex.
//
// - Main goroutine waits for all workers to finish with a WaitGroup.
func Wg_Mutex_Channel_Demo() {
	const workerCount = 3

	var (
		wg           sync.WaitGroup
		mu           sync.Mutex
		totalRevenue float64
		orderChan    = make(chan Order, 10)
	)

	// Launch worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for order := range orderChan {
				fmt.Printf("👷 Worker %d processing Order %d\n", workerID, order.ID)
				time.Sleep(time.Millisecond * 500) // Simulate processing delay

				mu.Lock()
				totalRevenue += order.Amount
				mu.Unlock()
			}
		}(i)
	}

	// Simulate order arrivals
	for i := 1; i <= 7; i++ {
		orderChan <- Order{
			ID:     i,
			Amount: float64(10 * i),
		}
	}
	close(orderChan) // No more orders

	// Wait for all workers to finish
	wg.Wait()

	fmt.Printf("\n✅ All orders processed. Total revenue: $%.2f\n", totalRevenue)
}
