package bufferedvsunbuffered

import (
	"fmt"
	"sync"
	"time"
)

func callAPI(id int) {
	fmt.Printf("🌐 API call %d at %v\n", id, time.Now().Format("15:04:05.000"))
	time.Sleep(200 * time.Millisecond) // simulate response time
}

func RateLimitedAPI() {
	var wg sync.WaitGroup
	limit := time.NewTicker(200 * time.Millisecond) // ~5 calls/sec
	defer limit.Stop()
	tasks := 15

	for i := 1; i <= tasks; i++ {
		wg.Add(1)
		<-limit.C // throttle
		go func(id int) {
			defer wg.Done()
			callAPI(id)
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ All API calls done")
}
