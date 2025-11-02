package advancedcontexts

import (
	"context"
	"fmt"
	"time"
)

func exampleTimeout(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Called the API")
	case <-ctxWithTimeout.Done():
		fmt.Println("Timeout!", ctxWithTimeout.Err())
	}
}

func Context_Timeout() {
	fmt.Println("\n==> Context Timeout")
	ctx := context.Background()
	exampleTimeout(ctx)
}
