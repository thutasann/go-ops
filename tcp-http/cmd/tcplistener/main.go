package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	lienter, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := lienter.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}
		for line := range getLinesChannel(conn) {
			fmt.Printf("read: %s\n", line)
		}
	}
}
