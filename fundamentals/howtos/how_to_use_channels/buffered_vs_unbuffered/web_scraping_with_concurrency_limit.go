package bufferedvsunbuffered

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fetchSite(id int, url string) {
	fmt.Printf("Worker %d fetching %s\n", id, url)
	time.Sleep(time.Duration(rand.Intn(200)+500) * time.Millisecond)
	fmt.Printf("Worker %d finished %s\n", id, url)
}

// Web Scraping with Concurrency Limit
// You scrape 1,000 websites.
// If you hit them all at once, you get blocked. Limit to 5 requests at a time.
func Web_Scraping_With_Concurrency_Limit() {
	rand.Seed(time.Now().UnixNano())
	sites := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://openai.com",
		"https://golang.org",
		"https://reddit.com",
		"https://news.ycombinator.com",
	}

	sem := make(chan struct{}, 3) // allow 3 concurrent scrapes
	var wg sync.WaitGroup

	for i, site := range sites {
		wg.Add(1)
		sem <- struct{}{} // acquire slot

		go func(id int, url string) {
			defer wg.Done()
			defer func() { <-sem }() // release slot

			fetchSite(id, url)
		}(i+1, site)
	}

	wg.Wait()
	fmt.Println("All sites fetched")
}
