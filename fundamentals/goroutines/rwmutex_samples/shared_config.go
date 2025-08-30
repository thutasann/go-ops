package rwmutexsamples

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Config) Get(key string) string {
	c.mu.RLock() // multiple reader can enter
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *Config) Set(key, value string) {
	c.mu.Lock() // exclusive write lock
	defer c.mu.Unlock()
	c.data[key] = value
}

func Shared_Config_Sample() {
	fmt.Print("\n===> Shared Config Sample")

	cfg := &Config{data: map[string]string{"mode": "dev"}}

	// Writer goroutine (updates config)
	go func() {
		time.Sleep(2 * time.Second)
		cfg.Set("mode", "prod")
		fmt.Println("Config updated!")
	}()

	// Read goroutines (read config)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				fmt.Printf("Reader %d sees mode: %s\n", id, cfg.Get("mode"))
				time.Sleep(1 * time.Second)
			}
		}(i)
	}

	wg.Wait()
}
