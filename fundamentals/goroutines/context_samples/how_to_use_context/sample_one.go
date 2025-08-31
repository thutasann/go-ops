package howtousecontext

import (
	"context"
	"fmt"
	"log"
	"time"
)

// How to Use Context Sample One
func SampleOne() {
	start := time.Now()
	fmt.Println("\n==> How to use Context Package Sample One")

	ctx := context.Background()
	userID := 10
	val, err := fetch_user_data(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
	fmt.Println("took: ", time.Since(start))
}

type UserResponse struct {
	value int
	err   error
}

// Fetch User Data
func fetch_user_data(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*700)
	defer cancel()
	respch := make(chan UserResponse)

	go func() {
		val, err := fetch_third_party_stuff_with_can_be_slow()
		respch <- UserResponse{value: val, err: err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took too long for %d", userID)
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

// Fetch Slow Third Party Stuff
func fetch_third_party_stuff_with_can_be_slow() (int, error) {
	time.Sleep(time.Millisecond * 500)
	return 666, nil
}
