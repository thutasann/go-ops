package syncvsasync

import (
	"fmt"
	"sync"
	"time"
)

func fetchProfile(id int) string {
	time.Sleep(2 * time.Second)
	return "Profile"
}

func fetchPosts(id int) string {
	time.Sleep(2 * time.Second)
	return "Posts"
}

func fetchFriends(id int) string {
	time.Sleep(2 * time.Second)
	return "Friends"
}

// 6.003175083s
func sync_microservices_calls() {
	fmt.Println("\n===> sync_microservices_calls")
	userID := 1

	start := time.Now()
	p := fetchProfile(userID)
	po := fetchPosts(userID)
	f := fetchFriends(userID)
	fmt.Println("Sync:", p, po, f, "took", time.Since(start))
}

// 2.001219542s
func async_microservices_callls() {
	fmt.Println("\n===> async_microservices_calls")
	userID := 1
	start := time.Now()
	var wg sync.WaitGroup
	var profile, posts, friends string

	wg.Add(3)
	go func() { profile = fetchProfile(userID); wg.Done() }()
	go func() { posts = fetchPosts(userID); wg.Done() }()
	go func() { friends = fetchFriends(userID); wg.Done() }()

	wg.Wait()
	fmt.Println("Async:", profile, posts, friends, "took", time.Since(start))
}

func MircoServices_Calls() {
	fmt.Println("\n===> MircoServices_Calls")
	sync_microservices_calls()
	async_microservices_callls()
}
