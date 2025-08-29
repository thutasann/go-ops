package channelsamples

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Producing:", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println("Consumed: ", val)
	}
}

func Producer_Consumer_Sample() {
	fmt.Println("\n====> Producer Consumer Sample")

	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
