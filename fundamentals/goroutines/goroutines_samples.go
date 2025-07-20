package goroutines

import (
	"fmt"
	"sync"
)

func GoRoutines_Sample_One() {
	fmt.Println("one")
	go testFunction()
	fmt.Println("two")
}

func testFunction() {
	fmt.Printf("three\n")
}

func GoRoutines_Sample_Two() {
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("one")
	go func() {
		defer wg.Done()
		testFunction()
	}()

	fmt.Println("two")
	wg.Wait()
}
