package goroutines

import (
	"fmt"
	"time"
)

func Channel_Sample_One() {
	fmt.Println("one")
	c := make(chan bool)
	go channelTestFunction(c)
	fmt.Println("two")
	areWeFinished := <-c
	fmt.Println("areWeFinished --> ", areWeFinished)
}

func channelTestFunction(c chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println("checking...")
		time.Sleep(1 * time.Second)
	}
	c <- true
}
