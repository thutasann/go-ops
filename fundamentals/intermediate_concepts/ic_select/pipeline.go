package icselect

import (
	"fmt"
	"time"
)

func generator(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			time.Sleep(200 * time.Millisecond)
			out <- n * n
		}
		close(out)
	}()

	return out
}

func printer(in <-chan int) {
	for n := range in {
		fmt.Println("📦 Output:", n)
	}
}

func Pipeline_Sample() {
	fmt.Println("\n==> Pipeline Sample")
	nums := generator([]int{1, 2, 3, 4, 5})
	squared := square(nums)
	printer(squared)
}
