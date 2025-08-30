package rwmutexsamples

import (
	"fmt"
	"sync"
	"time"
)

type Stock struct {
	mu     sync.RWMutex
	prices map[string]float64
}

func (s *Stock) GetPrice(symbol string) float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.prices[symbol]
}

func (s *Stock) UpdatePrice(symbol string, price float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.prices[symbol] = price
}

func Stock_Trading_Sample() {
	fmt.Print("\n===> Stock Trading Sample")

	stock := &Stock{prices: map[string]float64{"AAPL": 150.0}}

	// Writer: update price updates
	go func() {
		for i := 0; i < 5; i++ {
			newPrice := 150.0 + float64(i*5)
			stock.UpdatePrice("AAPL", newPrice)
			fmt.Println("Updated Price: ", newPrice)
			time.Sleep(2 * time.Second)
		}
	}()

	// Multiple readers: fetching live prices
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				price := stock.GetPrice("AAPL")
				fmt.Printf("Trader %d sees price: %.2f\n", id, price)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}

	wg.Wait()
}
