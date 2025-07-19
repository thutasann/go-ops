package main

import (
	"fmt"
	"os"
)

// go run *.go arg1 arg2
func Args_Sample() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usag: ./fundamentals <argument>\n")
		os.Exit(1)
	}

	fmt.Printf("hello world\nos.Args: %v\nArguments: %v\n", args, args[1:])
}
