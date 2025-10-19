package icchannels

import (
	"fmt"
	"time"
)

func chef(dishes []string, ch chan string) {
	for _, dish := range dishes {
		fmt.Printf("👨‍🍳 Cooking %s...\n", dish)
		time.Sleep(1 * time.Second)               // simulate work
		ch <- fmt.Sprintf("✅ %s is ready!", dish) // send dish to channel
	}
	close(ch)
}

func Chef_Waiter() {
	fmt.Println("\n===> Chef Waiter")

	ch := make(chan string)

	go chef([]string{"Pasta", "Soup", "Salad"}, ch)

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("🍽️ All dishes served! Kitchen closed.")
}
