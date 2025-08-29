package channelsamples

import (
	"fmt"
	"time"
)

func Select_With_Multiple_Channels() {
	fmt.Println("\n====> Select with multiple Channels")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Response from Service A"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Response from Service B"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Got:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Got:", msg2)
		case <-time.After(1500 * time.Millisecond):
			fmt.Println("Timeout! Skipping slow service...")
		}
	}
}
