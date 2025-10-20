package icselect

import (
	"context"
	"fmt"
	"time"
)

func Select_Context_Cancellation() {
	fmt.Println("\n===> Select Context Cancellation")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch <- "Done!"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("cancelled: ", ctx.Err())
	}
}
