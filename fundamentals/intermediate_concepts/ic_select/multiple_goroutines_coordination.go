package icselect

import (
	"fmt"
	"time"
)

func worker(name string, delay time.Duration, ch chan string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(delay)
		ch <- fmt.Sprintf("%s finished job %d", name, i)
	}
}

func Multiple_Goroutines_Coordination() {
	fmt.Println("\n==> Multiple Goroutines Coordination")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go worker("chef", 1*time.Second, ch1)
	go worker("waiter", 2*time.Second, ch2)

	for i := 0; i < 6; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}

	fmt.Println("All jobs done!")
}
