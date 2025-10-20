package icselect

import (
	"fmt"
	"time"
)

func Basic_Timeout_Pattern() {
	fmt.Println("\n==> Basic Timeout Pattern")

	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Operation Finished!"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(2 * time.Second):
		fmt.Println("⏰ Timeout! Operation took too long.")
	}
}
