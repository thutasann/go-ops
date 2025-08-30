package mutexsamples

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	s.m[key] = value
	s.mu.Unlock()
}

func (s *SafeMap) Get(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[key]
}

func Shared_Map_Sample() {
	fmt.Println("\n==> Shared Map Sample")

	safeMap := SafeMap{m: make(map[string]int)}

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key:%d", n), n)
		}(i)
	}

	wg.Wait()
	fmt.Println("Map contents:", safeMap.m)
}
