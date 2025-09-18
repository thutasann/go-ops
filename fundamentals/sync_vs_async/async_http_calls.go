package syncvsasync

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchAsync(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Fetched: ", url, "Status:", resp.StatusCode)
}

func Async_HTTP_Calls() {
	fmt.Println("\n==> Async_HTTP_Calls")
	var wg sync.WaitGroup
	urls := []string{"https://golang.org", "https://example.com", "https://github.com"}

	wg.Add(len(urls))
	for _, url := range urls {
		go fetchAsync(url, &wg)
	}
	wg.Wait()
}
