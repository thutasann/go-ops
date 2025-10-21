package icselect

import (
	"fmt"
	"math/rand"
	"time"
)

func mockAPICall() error {
	delay := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(delay)
	if delay > 2*time.Second {
		return fmt.Errorf("server timeout")
	}
	return nil
}

func Timeout_Retry() {
	fmt.Println("\n===> Timeout Retry")
	retries := 3
	for i := 1; i <= retries; i++ {
		done := make(chan error, 1)
		go func() {
			done <- mockAPICall()
		}()

		select {
		case err := <-done:
			if err != nil {
				fmt.Printf("⚠️ Attempt %d failed: %v\n", i, err)
				continue
			}
			fmt.Println("✅ API call succeeded!")
			return
		case <-time.After(2 * time.Second):
			fmt.Printf("⏰ Attempt %d timed out, retrying...\n", i)
		}
	}
	fmt.Println("❌ All retries failed.")
}
