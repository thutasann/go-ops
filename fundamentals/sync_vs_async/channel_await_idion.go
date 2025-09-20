package syncvsasync

import (
	"fmt"
	"time"
)

// channel-await idiom
func Await[T any](fn func() T) T {
	ch := make(chan T, 1)
	go func() { ch <- fn() }()
	return <-ch
}

func Channel_Await_Idiom() {
	fmt.Println("\n==> Channel Await Idiom")
	result := Await(func() string {
		time.Sleep(2 * time.Second)
		return "done!"
	})
	fmt.Println(result)
}
