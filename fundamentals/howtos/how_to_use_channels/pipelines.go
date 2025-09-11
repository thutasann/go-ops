package howtousechannels

import "fmt"

// gen -> sq -> out pattern
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func Pipeline_Small_Buffer() {
	fmt.Println("\n===> Pipeline Small Buffer")
	in := gen(2, 3, 4)
	out := sq(in)
	for v := range out {
		fmt.Println(v)
	}
}
