package main

import (
	"fmt"
	"net/http"
	"time"
)

// Asyn background task
func sendWelcomeEmail(email string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Background: welcome email sent to", email)
}

// Signup Handler
func signupHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}

	go sendWelcomeEmail(email)

	fmt.Fprintln(w, "Signed up with: ", email, "(email sent async)")
}

func startWebServer() {
	http.HandleFunc("/signup", signupHandler)
	fmt.Println("Web server on http://localhost:4000/signup?email=test@example.com")
	http.ListenAndServe(":4000", nil)
}
