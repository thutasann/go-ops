package howtousegoroutinesforaggdata

import (
	"fmt"
	"time"
)

func fifo_fetch_data(id int, out chan<- string) {
	time.Sleep(time.Duration(id) * time.Second)
	out <- fmt.Sprintf("Service %d data", id)
}

func fanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)
	for _, c := range channels {
		go func(c <-chan string) {
			for val := range c {
				out <- val
			}
		}(c)
	}
	return out
}

func FIFO_Pattern() {
	fmt.Println("\n===> FIFO Pattern")
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go fifo_fetch_data(1, c1)
	go fifo_fetch_data(2, c2)
	go fifo_fetch_data(3, c3)

	merged := fanIn(c1, c2, c3)

	for i := 0; i < 3; i++ {
		fmt.Println(<-merged)
	}
}
