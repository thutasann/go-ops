package wgsamples

import (
	"fmt"
	"sync"
	"time"
)

func wg_worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WG_Basic_Sample() {
	fmt.Print("\n====> WG Basic Sample")
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // tell WaitGrouop there's one more goroutine
		go wg_worker(i, &wg)
	}

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All Workers finished!")
}
