package syncvsasync

import (
	"fmt"
	"net/http"
	"time"
)

// Sync: blocks response until email is "sent"
func signupSync(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintln(w, "User signed up & email sent!")
}

// Async: responds immediately, sends email in background
func signupAsync(w http.ResponseWriter, r *http.Request) {
	go func() {
		time.Sleep(3 * time.Second) // simulate email sending
		fmt.Println("Background: welcome email sent")
	}()
	fmt.Fprintln(w, "User signed up! (email sending async)")
}

func WebServer_Sync_Vs_Async() {
	fmt.Println("\n===> WebServer_Sync_Vs_Async")
	http.HandleFunc("/sync", signupSync)
	http.HandleFunc("/async", signupAsync)
	fmt.Println("Server running on http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
