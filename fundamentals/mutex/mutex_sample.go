package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type mytype struct {
	counter int
	mu      sync.Mutex
}

func (m *mytype) IncreaseCounter() {
	m.mu.Lock()
	m.counter++
	m.mu.Unlock()
}

func (m *mytype) GetCounter() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.counter
}

func Mutex_Sample() {
	var wg sync.WaitGroup
	myTypeInstance := mytype{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(m *mytype) {
			defer wg.Done()

			m.IncreaseCounter()
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

			if m.GetCounter() == 5 {
				fmt.Println("Found counter == 5")
			}
		}(&myTypeInstance)
	}

	wg.Wait() // wait for all goroutines
	fmt.Printf("Final counter: %d\n", myTypeInstance.GetCounter())
}
