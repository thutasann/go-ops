package wgsamples

import (
	"fmt"
	"sync"
	"time"
)

// A WaitGroup is used to wait for a collection of goroutines to finish. It acts as a counter:
//
// Imagine you're building a service that fetches product data from multiple microservices concurrently: pricing, inventory, reviews. You want to wait until all data is fetched before combining the result.
//
// - Add(n) — Adds n goroutines to wait for.
// - Done() — Each goroutine calls this when it's finished.
// - Wait() — Blocks until the counter goes back to zero.

func fetchPrice(wg *sync.WaitGroup, result *string) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	*result = "$19.99"
	fmt.Println("Fetched price")
}

func fetchInventory(wg *sync.WaitGroup, result *string) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	*result = "24 items"
	fmt.Println("Fetched inventory")
}

func fetchReviews(wg *sync.WaitGroup, result *string) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	*result = "⭐️⭐️⭐️⭐️☆"
	fmt.Println("Fetched reviews")
}

func WaitGroup_Sample() {
	var wg sync.WaitGroup

	var price, inventory, reviews string

	wg.Add(3) // waiting for 3 tasks

	go fetchPrice(&wg, &price)
	go fetchInventory(&wg, &inventory)
	go fetchReviews(&wg, &reviews)

	wg.Wait() // Wait until all are done

	fmt.Println("\n📦 Product Summary:")
	fmt.Println("Price:", price)
	fmt.Println("Inventory:", inventory)
	fmt.Println("Reviews:", reviews)
}
