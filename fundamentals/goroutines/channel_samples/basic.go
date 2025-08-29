package channelsamples

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Worker done!"
}

func Basic_Channel_Sample() {
	fmt.Println("\n====> Basic Channel Sample")
	ch := make(chan string)

	go worker(ch)

	fmt.Println("Waiting for worker...")

	msg := <-ch
	fmt.Println("Received: ", msg)
}

func Buffer_Channel_Sample() {
	fmt.Println("\n====> Buffer Channel Sample")

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 <-- will get error

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
