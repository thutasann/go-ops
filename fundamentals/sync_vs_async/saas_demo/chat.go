package main

import (
	"bufio"
	"fmt"
	"net"
)

var chatClients = make(map[net.Conn]bool)
var chatMessages = make(chan string)

func handleChatConn(conn net.Conn) {
	chatClients[conn] = true
	scanner := bufio.NewScanner(conn)

	// Read async
	go func() {
		for scanner.Scan() {
			msg := scanner.Text()
			chatMessages <- msg
		}
		delete(chatClients, conn)
		conn.Close()
	}()
}

func broadcaster() {
	for msg := range chatMessages {
		for c := range chatClients {
			fmt.Fprintln(c, msg)
		}
	}
}

func startChatServer() {
	ln, _ := net.Listen("tcp", ":4001")
	go broadcaster()

	fmt.Println("Chat server started on :4001 (use telnet localhost 4001)")
	for {
		conn, _ := ln.Accept()
		go handleChatConn(conn)
	}
}
