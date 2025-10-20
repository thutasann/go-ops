package icselect

import "fmt"

func Non_Blocking_Send_Receive() {
	fmt.Println("\n===> Non Blocking Send & Receive")

	ch := make(chan string, 1)

	select {
	case ch <- "Hello":
		fmt.Println("✅ Message sent")
	default:
		fmt.Println("Channel full, message dropped")
	}

	select {
	case msg := <-ch:
		fmt.Println("Received: ", msg)
	default:
		fmt.Println("No message available")
	}
}
