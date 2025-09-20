package syncvsasync

import (
	"fmt"
	"sync"
)

func Cond_Sample() {
	fmt.Println("\n==> Cond Sample")

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	// waiter goroutine
	go func() {
		cond.L.Lock()
		for !ready {
			cond.Wait() // wait until signaled
		}
		fmt.Println("got the signal!")
		cond.L.Unlock()
	}()

	// Signaler
	cond.L.Lock()
	ready = true
	cond.Signal()
	cond.L.Unlock()
}
