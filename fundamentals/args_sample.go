package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// go run *.go arg1 arg2
func args_sample() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usag: ./fundamentals <argument>\n")
		os.Exit(1)
	}

	fmt.Printf("hello world\nos.Args: %v\nArguments: %v\n", args, args[1:])

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in an invalid format: %s\n", err)
		os.Exit(1)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP Status Code: %d\nBody: %s\n", response.StatusCode, body)
}
