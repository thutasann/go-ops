package howtoslices

import "fmt"

// ============= STACK SAMPLE =============

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	n := len(*s)
	x := (*s)[n-1]
	*s = (*s)[:n-1]
	return x
}

func Stack_Slice_Sample() {
	fmt.Println("\n==> Stack Slice Sample")
	var st Stack
	st.Push(10)
	st.Push(20)
	fmt.Println(st.Pop())
}

// ============= QUEUE SAMPLE =============

type Queue []int

func (q *Queue) Enqueue(x int) {
	*q = append(*q, x)
}

func (q *Queue) Dequeue() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func Queue_Sample() {
	fmt.Println("\n===> Queue Sample")
	var queue Queue
	queue.Enqueue(10)
	queue.Enqueue(20)
	fmt.Println(queue.Dequeue())
}
