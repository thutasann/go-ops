package bufferedvsunbuffered

import (
	"fmt"
)

// Basic send/receive sync
func Basic_Send_Receive_Sync() {
	fmt.Println("\n===> Basic Send / Receive")
	ch := make(chan int)

	go func() {
		fmt.Println("sending...")
		ch <- 42 // blocks until someone receives
		fmt.Println("done sending")
	}()

	fmt.Println("receiving...")
	v := <-ch // unblocks sender
	fmt.Println("got", v)
}

// Deadlock if no receiver
func Deadlock_If_No_Receiver() {
	fmt.Println("\n===> Deadlock if not receiver")
	ch := make(chan int)
	ch <- 99 // blocks forever, no one to receive
}

// Ping Pong
func Ping_Pong() {
	fmt.Println("\n===> Ping Pong")
	ch := make(chan string)

	go func() {
		msg := <-ch
		fmt.Println("worker got: ", msg)
		ch <- "pong"
	}()

	ch <- "ping" // blocks until worker receives
	fmt.Println("maing sent ping")
	reply := <-ch
	fmt.Println("main got: ", reply)
}

// Buffered channel examples (Fill buffer then block)
func Fill_Buffer_Then_Block() {
	fmt.Println("\n===> Fill buffer then block")

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("buffer full")

	go func() {
		fmt.Println("trying to send 3...")
		ch <- 3 // blocks here until main receives
		fmt.Println("sent 3")
	}()

	fmt.Println("main receives (frees space):", <-ch) // frees space
	fmt.Println("main receives:", <-ch)
	fmt.Println("main receives:", <-ch)
}
