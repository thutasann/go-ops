package howtousegoroutinesforaggdata

import (
	"fmt"
	"strings"
	"sync"
)

func Word_Frequency_Counter() {
	fmt.Println("\n==> Word Frequency Counter")

	texts := []string{
		"hello world hello",
		"go is fun go fast",
		"goroutines are awesome",
	}

	wordCount := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, t := range texts {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			words := strings.Fields(s)
			local := make(map[string]int)
			for _, w := range words {
				local[w]++
			}

			// Aggregate into shared map (mutex protected)
			mu.Lock()
			for k, v := range local {
				wordCount[k] += v
			}
			mu.Unlock()
		}(t)
	}

	wg.Wait()
	fmt.Println("Word frequencies:", wordCount)
}
