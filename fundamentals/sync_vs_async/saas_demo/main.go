package main

import (
	"fmt"
	"time"
)

// SAAS Demo about Async
// http://localhost:4000/signup?email=you@example.com
// telnet localhost 4001
func main() {
	// Start web server
	go startWebServer()

	// Start worker pool
	go startWorkerPool()

	// Start Scheduler
	go startScheduler()

	// Start chat
	go startChatServer()

	fmt.Println("SaaS Demo running...")
	for {
		time.Sleep(1 * time.Hour) // keep alive
	}
}
