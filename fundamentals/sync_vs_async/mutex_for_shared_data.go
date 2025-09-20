package syncvsasync

import (
	"fmt"
	"sync"
)

// Prevents race conditions on counter.
func Mutex_For_Shared_Data() {
	fmt.Println("\n===> Mutex For Shared Data")
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}
