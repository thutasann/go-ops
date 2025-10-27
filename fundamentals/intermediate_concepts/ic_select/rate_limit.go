package icselect

import (
	"fmt"
	"time"
)

func RateLimit() {
	fmt.Println("\n==> Rate Limit")
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Allow 1 request every 500ms
	limiter := time.Tick(500 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("Processing request", req, "at", time.Now().Format("15:04:05"))
	}
}
