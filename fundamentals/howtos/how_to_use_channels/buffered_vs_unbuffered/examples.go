package bufferedvsunbuffered

import (
	"fmt"
	"time"
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

// Drain Empty Buffer blocks
func Drain_Empty_Buffer_Blocks() {
	fmt.Println("\n===> Drain Empty Buffer Blocks")
	ch := make(chan string, 1)

	go func() {
		fmt.Println("waiting to receive...")
		fmt.Println("got: ", <-ch) // blocks until main sends
	}()

	ch <- "hello"
}

// Producer Faster than consumer
// Producer can run ahead until buffer fills, then it must wait.
func Producer_Faster_Than_Consumer() {
	ch := make(chan int, 3)

	// producer
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("[producer] sending:", i)
			ch <- i
			fmt.Println("[producer] sent:", i)
		}
		close(ch)
	}()

	// consumer
	for v := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("[consumer] got:", v)
	}
}

// Worker Pool Concurrency Limit
func Worker_Pool_Concurrency_Limit() {
	jobs := 5
	sem := make(chan struct{}, 2) // max 2 workers at a time

	for i := 1; i <= jobs; i++ {
		sem <- struct{}{}
		go func(id int) {
			defer func() { <-sem }()
			fmt.Println("working:", id)
			time.Sleep(500 * time.Millisecond)
			fmt.Println("done:", id)
		}(i)
	}

	// wait for all the slots to be freed
	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
}
