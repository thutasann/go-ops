package syncvsasync

import (
	"bufio"
	"fmt"
	"net"
)

var clients = make(map[net.Conn]bool)
var messages = make(chan string)

func handleConn(conn net.Conn) {
	clients[conn] = true
	scanner := bufio.NewScanner(conn)

	// Read async
	go func() {
		for scanner.Scan() {
			msg := scanner.Text()
			messages <- msg
		}
		delete(clients, conn)
		conn.Close()
	}()
}

func broadcaster() {
	for msg := range messages {
		for c := range clients {
			fmt.Fprintln(c, msg)
		}
	}
}

// telnet localhost 8081
func Async_Chat_Server() {
	fmt.Println("\n===> Async_Chat_Server")
	ln, _ := net.Listen("tcp", ":8081")
	go broadcaster()

	fmt.Println("Chat server started on :8081")
	for {
		conn, _ := ln.Accept()
		go handleConn(conn)
	}
}
