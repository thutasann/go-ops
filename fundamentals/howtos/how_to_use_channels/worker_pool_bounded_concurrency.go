package howtousechannels

import (
	"fmt"
	"sync"
)

func bounded_concurrency_worker(id int, jobs <-chan int, wg *sync.WaitGroup, sem chan struct{}) {
	fmt.Println("\n===> Bounded Concurrency Worker")
}
