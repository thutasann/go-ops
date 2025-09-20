package syncvsasync

import (
	"fmt"
	"sync"
)

func Once_Sample() {
	fmt.Println("\n==> Once Sample")

	var once sync.Once

	for i := 0; i <= 3; i++ {
		go once.Do(func() {
			fmt.Println("Runs only once")
		})
	}

	select {}
}
