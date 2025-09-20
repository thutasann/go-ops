package main

import (
	"fmt"
	"time"
)

func generateReport() {
	fmt.Println("[Scheduler] Generating report at: ", time.Now())
	time.Sleep(2 * time.Second) // heavy work
	fmt.Println("[Scheduler] Report ready at: ", time.Now())
}

func startScheduler() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			go generateReport()
		}
	}()
}
