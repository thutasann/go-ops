package contextsamples

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func context_fetch_data(context context.Context, url string) ([]byte, error) {
	req, _ := http.NewRequestWithContext(context, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Context_Fetch_Data_Sample() {
	fmt.Println("\n====> Context Fetch Data Sample")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	data, err := context_fetch_data(ctx, "https://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(data))
}
