package syncvsasync

import (
	"fmt"
	"net/http"
)

func fetchSync(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Fetched: ", url, "Status:", resp.StatusCode)
}

func Sync_HTTP_Calls_Sample() {
	fmt.Println("\n==> Sync_HTTP_Calls_Sample")
	fetchSync("https://golang.org")
	fetchSync("https://example.com") // runs only after the first finishes
}
