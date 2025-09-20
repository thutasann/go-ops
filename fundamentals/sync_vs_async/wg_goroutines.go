package syncvsasync

import (
	"fmt"
	"sync"
	"time"
)

func wg_worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[wg worker] Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("[wg worker] Worker %d done\n", id)
}

// In Go, synchronization means coordinating the execution of multiple goroutines so they don’t:
// - Race over shared data, or
// - Exit too early before others finish.
// Since goroutines run concurrently, you need sync mechanisms to:
// - Ensure correct access to shared resources.
// - Wait for tasks to complete.
// - Control execution order when needed.
func Waiting_For_Goroutines() {
	fmt.Println("\n==> Waiting for Goroutines")
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go wg_worker(w, &wg)
	}

	wg.Wait() // waits until all goroutines call Done()
	fmt.Println("All Workers finished")
}
