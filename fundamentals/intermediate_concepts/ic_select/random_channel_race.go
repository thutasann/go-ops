package icselect

import (
	"fmt"
	"math/rand"
	"time"
)

func racer(name string, ch chan string) {
	delay := time.Duration(rand.Intn(2000)+1000) * time.Millisecond
	time.Sleep(delay)
	ch <- fmt.Sprintf("%s finished in %v", name, delay)
}

func Random_Channel_Race() {
	fmt.Println("\n==> Random Channel Race")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go racer("🐢 Turtle", ch1)
	go racer("🐇 Rabbit", ch2)

	select {
	case msg := <-ch1:
		fmt.Println("Winner:", msg)
	case msg := <-ch2:
		fmt.Println("Winner:", msg)
	}
}
