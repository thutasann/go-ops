package advancedcontexts

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("API response!")
	case <-ctx.Done():
		fmt.Println("Oh no the context expired")
		http.Error(w, "Request context timeout!", http.StatusRequestTimeout)
		return
	}
}

func Context_HTTP() {
	fmt.Println("\n==> HTTP Context")
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
