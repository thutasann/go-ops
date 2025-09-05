package howtousegoroutinesforaggdata

import (
	"fmt"
	"sync"
	"time"
)

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "Bob"
}

func fetchUserLikes(userName string, repch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	repch <- 11
	wg.Done()
}

func fetchUserMatch(userName string, repch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	repch <- "ANNA"
	wg.Done()
}

func SampleOne() {
	fmt.Println("\n==> Sample One")

	start := time.Now()
	userName := fetchUser()
	respch := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)

	wg.Wait() // block until 2 wg.Done()
	close(respch)

	for resp := range respch {
		fmt.Print("resp:", resp)
	}

	fmt.Println("took:", time.Since(start))

}
