package net_cat

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func Netcat() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8989"
	}
	_, err := strconv.Atoi(port)
	if err != nil || len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	room := &chatRoom{clients: make([]client, 0)}
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Println("Server started, listening on port", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(room, conn)
	}
}
