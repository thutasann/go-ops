package icchannels

import (
	"fmt"
	"time"
)

// chef sends the dishes to the channel
func cl_chef(dishes []string, ch chan string) {
	for _, dish := range dishes {
		fmt.Printf("👨‍🍳 Cooking %s...\n", dish)
		time.Sleep(1 * time.Second)
		ch <- fmt.Sprintf("✅ %s is ready!", dish)
	}
	close(ch)
}

// waiter receives the dishes from the channel
func cl_waiter(ch chan string, done chan bool) {
	for dish := range ch {
		fmt.Printf("🧑‍🍽️ Serving dish... %s\n", dish)
		time.Sleep(500 * time.Millisecond)
	}
	done <- true // signal that serving is finished
}

// Challenge Chef Waiter
func Challenge_Chef_Waiter() {
	fmt.Println("\n==> Challenge Chef Waiter")

	// buffered channel (capcaity = 2)
	ch := make(chan string, 2)

	// separate channel to signal completion
	done := make(chan bool)

	// Start chef and waiter concurrently
	go cl_chef([]string{"Pasta", "Soup", "Salad", "Steak"}, ch)
	go cl_waiter(ch, done)

	<-done
	fmt.Println("🎉 All dishes cooked and served successfully!")
}
