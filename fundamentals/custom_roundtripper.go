package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// CustomTransport wraps an existing RoundTripper (usually http.DefaultTransport)
type CustomTransport struct {
	Transport http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface
func (ct *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()

	fmt.Printf("Request: %s %s\n", req.Method, req.URL)

	resp, err := ct.Transport.RoundTrip(req)
	if err != nil {
		fmt.Printf("Error during request: %v\n", err)
		return nil, err
	}

	duration := time.Since(start)
	fmt.Printf("📥 Response: %d %s (Time: %s)\n", resp.StatusCode, http.StatusText(resp.StatusCode), duration)

	return resp, nil
}

func Custom_RoundTrip_Sample() {
	client := &http.Client{
		Transport: &CustomTransport{Transport: http.DefaultTransport},
		Timeout:   10 * time.Second,
	}

	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("🔽 Body:", string(body))
}
