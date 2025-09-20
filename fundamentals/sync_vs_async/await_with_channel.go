package syncvsasync

import (
	"fmt"
	"time"
)

// fetch data
func await_fetch_data() <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from async work"
	}()
	return ch
}

// Here result := <-await_fetch_data() works like
// await await_fetch_data() → blocks this function only, not the whole program.
func Await_With_Channel() {
	fmt.Println("\n==> Await_With_Channel")
	results := <-await_fetch_data()
	fmt.Println("results :>> ", results)
}
