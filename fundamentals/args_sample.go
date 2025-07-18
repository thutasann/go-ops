package main

import (
	"fmt"
	"os"
)

// go run main.go arg1 arg2
func args_sample() {
	var args []string
	args = os.Args
	fmt.Printf("hello world\nos.Args: %v\nArguments: %v\n", args, args[1:])
}
