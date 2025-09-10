package howtousechannels

import (
	"fmt"
	"time"
)

func Basic_Sample_One() {
	fmt.Println("\n==> Channels Basic Sample One")

	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	val := <-ch
	fmt.Println("Got:", val)
}

func Buffered_Vs_Unbuffered() {
	fmt.Println("\n===> Buffered_Vs_Unbuffered")

	unbuffered := make(chan string)  // blocks until received
	buffered := make(chan string, 2) // can hold 2 items without waiting

	go func() {
		unbuffered <- "One"
	}()
	fmt.Println("unbuffered: ", <-unbuffered)

	buffered <- "two"
	buffered <- "three"

	fmt.Println("buffered: ", <-buffered)
	fmt.Println("buffered: ", <-buffered)
}

func Unbuffered_Channel_HandShake() {
	fmt.Println("\n==> Unbuffered_Channel_HandShake")

	ch := make(chan string) // unbuffered

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "ready" // blocks until receiver reads
		fmt.Println("sent")
	}()

	msg := <-ch
	fmt.Println("received:", msg)
}

func Buffered_Channel_Decoupling() {
	fmt.Println("\n==> Buffered_Channel_Decoupling")
	ch := make(chan int, 3)

	for i := 0; i < 3; i++ {
		ch <- i // never blocks (fits in buffer)
		fmt.Println("pushed: ", i)
	}

	go func() {
		time.Sleep(300 * time.Millisecond)
		for v := range ch {
			fmt.Println("got: ", v)
		}
	}()

	// This send will block when buffer full:
	fmt.Println("sending 99 (will block if buffer full)")
	ch <- 99
	close(ch)
	time.Sleep(500 * time.Millisecond)
}

func Non_Blocking_Send_Receive() {
	fmt.Println("\n==> Non blocking send / receive")

	ch := make(chan int, 1)
	select {
	case ch <- 1:
		fmt.Println("sent 1")
	default:
		fmt.Println("would block, skip send")
	}

	select {
	case v := <-ch:
		fmt.Println("received: ", v)
	default:
		fmt.Println("would block on receive")
	}
}
