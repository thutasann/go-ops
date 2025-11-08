package advancedchannels

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Messages struct {
	chats   []string
	friends []string
}

func Fetch_User_By_Name() {
	now := time.Now()
	id := getUserByName("John")
	println("userId: ", id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Messages, 2)

	wg.Add(2)

	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}

	log.Println(time.Since(now))
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s-2", name)
}

func getUserFriends(id string, ch chan<- *Messages, wg *sync.WaitGroup) {
	ch <- &Messages{
		friends: []string{
			"john",
			"jane",
			"joe",
			"james",
		},
	}
	wg.Done()
}

func getUserChats(id string, ch chan<- *Messages, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	ch <- &Messages{
		chats: []string{
			"John",
			"Jane",
			"Joe",
		},
	}

	wg.Done()
}

func leaky() {
	ch := make(chan int)

	go func() {
		msg := <-ch
		log.Println(msg)
	}()
}
