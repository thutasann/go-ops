package bufferedvsunbuffered

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DownloadTask struct {
	ID    int
	URL   string
	Retry int
}

func downloadFile(task DownloadTask) error {
	if rand.Float32() < 0.3 {
		return fmt.Errorf("network error for %s", task.URL)
	}
	time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
	fmt.Printf("✅ Downloaded file %d from %s\n", task.ID, task.URL)
	return nil
}

func FileDownloader() {
	fmt.Println("\n===> File Downloader")

	rand.Seed(time.Now().UnixNano())

	tasks := []DownloadTask{
		{1, "https://file1.com", 2},
		{2, "https://file2.com", 2},
		{3, "https://file3.com", 2},
		{4, "https://file4.com", 2},
		{5, "https://file5.com", 2},
		{6, "https://file6.com", 2},
	}

	sem := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		sem <- struct{}{}

		go func(t DownloadTask) {
			defer wg.Done()
			defer func() { <-sem }()

			for attempts := 1; attempts <= t.Retry+1; attempts++ {
				err := downloadFile(t)
				if err != nil {
					fmt.Printf("⚠️  Failed to download %s (attempt %d)\n", t.URL, attempts)
					time.Sleep(500 * time.Millisecond)
					continue
				}
				break
			}
		}(task)
	}

	wg.Wait()
	fmt.Println("🎉 All downloads done")
}
