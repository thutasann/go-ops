package syncvsasync

import (
	"fmt"
	"time"
)

// async so scheduler keeps ticking
func generateReport() {
	fmt.Println("Generating report...", time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println("Report ready:", time.Now())
}

func Cron_Like_Scheduler() {
	fmt.Println("\n===> Cron_Like_Scheduler")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		go generateReport()
	}
}
