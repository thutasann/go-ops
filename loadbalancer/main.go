package main

import (
	"fmt"
	"net/http"
)

// Simple LoadBalancer in GoLang
func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com/"),
		newSimpleServer("https://www.bing.com/"),
		newSimpleServer("https://www.duckduckgo.com/"),
	}

	lb := NewLoadBalancer("8000", servers)

	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		lb.serveProxy(rw, r)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Serving requests at 'localhost:%s'\n", lb.port)

	http.ListenAndServe(":"+lb.port, nil)
}
