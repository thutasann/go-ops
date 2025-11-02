package advancedcontexts

import (
	"context"
	"fmt"
)

func exampleWithValues() {
	type key int
	const userKey key = 0

	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, userKey, "123")

	if userID, ok := ctxWithValue.Value(userKey).(string); ok {
		fmt.Println("this is the user id", userID)
	} else {
		fmt.Println("this is a protected route - no userID found")
	}
}

func ContextWithValue() {
	fmt.Println("\n==> Context With Value")
	exampleWithValues()
}
