package howtoslices

import "fmt"

func ringBuffer(size int) func(msg string) []string {
	buffer := make([]string, 0, size)
	return func(msg string) []string {
		if len(buffer) == size {
			buffer = buffer[1:]
		}
		buffer = append(buffer, msg)
		return buffer
	}
}

func Ring_Buffer_Chat_Messages() {
	fmt.Println("\n==> Ring Buffer Chat Messages")
	addMsg := ringBuffer(3)
	fmt.Println(addMsg("hi"))
	fmt.Println(addMsg("hello"))
	fmt.Println(addMsg("hey"))
	fmt.Println(addMsg("hola")) // oldest ("hi") dropped
}
