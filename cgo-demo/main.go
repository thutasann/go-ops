package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "It's working")
}

// CGO Demo
//
// - GOOS=darwin ARCH=amd64 go build -o main-darwin64 main.go
//
// - CGO_ENABLED=0 go build -o main-nocgo main.go
func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
