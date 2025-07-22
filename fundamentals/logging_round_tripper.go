package main

import (
	"log"
	"net/http"
	"time"
)

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := lrt.Proxied.RoundTrip(req)
	log.Printf("Request to %s took %v", req.URL, time.Since(start))
	return resp, err
}

func Loggin_Round_Tripper() {
	client := &http.Client{
		Transport: LoggingRoundTripper{Proxied: http.DefaultTransport},
	}
	client.Get("https://httpbin.org/get")
}
