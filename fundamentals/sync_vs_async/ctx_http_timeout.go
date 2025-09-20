package syncvsasync

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func context_fetch_data(ctx context.Context) (string, error) {
	select {
	case <-time.After(3 * time.Second): // simulating slow work
		return "Data ready", nil
	case <-ctx.Done(): // context cancelled or timed out
		return "", ctx.Err()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// timeout after 2
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	data, err := context_fetch_data(ctx)
	if err != nil {
		http.Error(w, "Request canceled: "+err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, data)
}

func Context_With_Timeout_HTTP() {
	fmt.Println("\n==> Context With Timeout")
	http.HandleFunc("/", handler)
	fmt.Println("Server on http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
