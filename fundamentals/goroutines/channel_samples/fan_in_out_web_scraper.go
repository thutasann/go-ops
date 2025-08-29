package channelsamples

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fetch(url string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	ch <- fmt.Sprintf("Fetched: %s", url)
}

func Fan_In_Out_Web_Scraper_Sample() {
	fmt.Println("\n====> Buffer Channel Sample")

	urls := []string{"google.com", "facebook.com", "twitter.com"}
	ch := make(chan string)
	var wg sync.WaitGroup

	// Fan-out: multiple workers fetch concurrently
	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg, ch)
	}

	// Fan-out: collect results in separate goroutine
	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
}
