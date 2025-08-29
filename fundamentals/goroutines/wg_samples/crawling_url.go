package wgsamples

import (
	"fmt"
	"net/http"
	"sync"
)

func crawl_fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erorr fetching:", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(url, "-> ", resp.Status)
}

func Crawling_Sample() {
	fmt.Println("\n===> Crawling Sample")

	var wg sync.WaitGroup

	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://google.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go crawl_fetch(url, &wg)
	}

	wg.Wait()
	fmt.Println("All URLs fetched")
}
