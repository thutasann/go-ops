package howtousegoroutinesforaggdata

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func context_fetch(ctx context.Context, id int, ch chan<- string) {
	delay := time.Duration(rand.Intn(5)) * time.Second
	select {
	case <-time.After(delay):
		ch <- fmt.Sprintf("Worker %d finished in %v", id, delay)
	case <-ctx.Done():
		fmt.Printf("Worker %d cancelled\n", id)
	}
}

func Context_Aware_Aggr() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch := make(chan string, 3)

	for i := 1; i <= 3; i++ {
		go context_fetch(ctx, i, ch)
	}

	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch:
			fmt.Println("✅", msg)
		case <-ctx.Done():
			fmt.Println("⏰ Timeout reached, cancelling aggregation")
			return
		}
	}
}
