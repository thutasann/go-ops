package syncvsasync

import (
	"fmt"
	"sync"
	"time"
)

func wg_fetch_data(wg *sync.WaitGroup, results *string) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	*results = "Hello from Async Work"
}

func Await_With_WG() {
	fmt.Println("\n==> Await With WG")
	var wg sync.WaitGroup
	var res string

	wg.Add(1)
	go wg_fetch_data(&wg, &res)

	wg.Wait() // wait like "await"
	fmt.Println(res)
}
