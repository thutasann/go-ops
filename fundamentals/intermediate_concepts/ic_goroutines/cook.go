package icgoroutines

import (
	"fmt"
	"time"
)

func cook(dish string, duration time.Duration) {
	fmt.Printf("👨‍🍳 Starting to cook %s...\n", dish)
	time.Sleep(duration)
	fmt.Printf("✅ %s is ready!\n", dish)
}

func Cooking() {
	fmt.Println("\n===> IC Goroutines")

	go cook("Pasta", 3*time.Second)
	go cook("Soup", 2*time.Second)
	go cook("Salad", 1*time.Second)

	time.Sleep(4 * time.Second)

	fmt.Println("🎉 All dishes are ready! Bon appétit!")
}
