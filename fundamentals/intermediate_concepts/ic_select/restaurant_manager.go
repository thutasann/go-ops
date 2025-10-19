package icselect

import (
	"fmt"
	"time"
)

// Simulate a chef sending dishes
func chef(ch chan string) {
	dishes := []string{"Pasta", "Soup", "Salad"}
	for _, dish := range dishes {
		time.Sleep(time.Duration(1+len(dish)%2) * time.Second)
		ch <- fmt.Sprintf("👨‍🍳 %s is ready!", dish)
	}
	close(ch)
}

// Simulate a waiter sending updates
func waiter(ch chan string) {
	status := []string{"Taking order", "Serving drinks", "Delivering food"}
	for _, s := range status {
		time.Sleep(2 * time.Second)
		ch <- fmt.Sprintf("🧑‍🍽️ %s", s)
	}
	close(ch)
}

func Restaurant_Manager() {
	fmt.Println("\n===> Restaurant Manager")

	chefCh := make(chan string)
	waiterCh := make(chan string)

	go chef(chefCh)
	go waiter(waiterCh)

	timeout := time.After(6 * time.Second)

	for {
		select {
		case msg, ok := <-chefCh:
			if !ok {
				chefCh = nil // disable closed channel
			} else {
				fmt.Println(msg)
			}

		case msg, ok := <-waiterCh:
			if !ok {
				waiterCh = nil
			} else {
				fmt.Println(msg)
			}

		case <-timeout:
			fmt.Println("⏰ Timeout! Kitchen closed early.")
			return
		}

		// Exit when both channels are closed
		if chefCh == nil && waiterCh == nil {
			break
		}
	}

	fmt.Println("🎉 All tasks done successfully!")
}
