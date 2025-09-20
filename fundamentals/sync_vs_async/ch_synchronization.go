package syncvsasync

import (
	"fmt"
	"time"
)

func Synchronization_With_Channel() {
	fmt.Println("\n==> Synchronization with channel")

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from Goroutine"
	}()

	result := <-ch
	fmt.Println("result: ", result)
}
