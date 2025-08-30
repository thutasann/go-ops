package mutexsamples

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func Counter_Sample_Two() {
	fmt.Println("\n===> Counter Sample")
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Final counter: ", counter)
}
