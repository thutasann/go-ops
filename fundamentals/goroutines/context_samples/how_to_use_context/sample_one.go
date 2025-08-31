package howtousecontext

import (
	"context"
	"fmt"
	"log"
	"time"
)

// With Background Sample
func With_Background_Sample() {
	start := time.Now()
	fmt.Println("\n==> With Background Sample")

	ctx := context.Background()
	userID := 10
	val, err := fetch_user_data(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
	fmt.Println("took: ", time.Since(start))
}

type contextKey string

const fooKey contextKey = "foo"

// With Value Sample
func With_Value_Sample() {
	start := time.Now()
	fmt.Println("\n==> With Value Sample")

	ctx := context.WithValue(context.Background(), fooKey, "bar")

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
	val_from_context := ctx.Value(fooKey)
	fmt.Println("value from context: ", val_from_context)

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
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
	time.Sleep(time.Millisecond * 100)
	return 666, nil
}
