package net_cat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(room *chatRoom, conn net.Conn) {
	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	logo := Logo()
	conn.Write([]byte(logo))
	for {
		conn.Write([]byte("[ENTER YOUR NAME]:"))
		reader := bufio.NewScanner(conn)
		reader.Scan()
		name := strings.TrimSpace(reader.Text())
		if len(name) == 0 {
			conn.Write([]byte("Name can only contain alphabets and digits\n"))
			continue
		}
		for _, c := range name {
			if !NameCheck(c) {
				conn.Write([]byte("Name can only contain alphabets and digits\n"))
				continue
			}
		}
		room.mutex.Lock()
		taken := false
		for _, c := range room.clients {
			if c.name == name {
				taken = true
				break
			}
		}
		if taken {
			room.mutex.Unlock()
			conn.Write([]byte(fmt.Sprintf("Sorry, the name %s is already taken. Please choose a different name.\n", name)))
			continue
		}
		if len(room.clients) >= maxClients {
			room.mutex.Unlock()
			conn.Write([]byte("Sorry, maximum number of clients reached\n"))
			return
		}
		room.clients = append(room.clients, client{conn, name})
		name = strings.TrimSuffix(name, "\n")
		clientIdx := len(room.clients) - 1
		room.mutex.Unlock()
		conn.Write([]byte(fmt.Sprintf("[%s][SERVER]: You have joined the chat\n", timestamp)))
		room.mutex.Lock()
		for i, c := range room.clients {
			if i != clientIdx {
				c.conn.Write([]byte(fmt.Sprintf("\n[%s][SERVER]: %s has joined the chat\n", timestamp, name)))
			}
		}
		room.mutex.Unlock()
		room.mutex.Lock()
		if len(room.history) > 0 {
			// conn.Write([]byte("Chat history:\n"))
			for _, msg := range room.history {
				conn.Write([]byte(msg))
			}
			conn.Write([]byte("\n"))
		}
		room.mutex.Unlock()
		go handleMessages(room, clientIdx)
		break
	}
}

func Logo() string {
	penguin, err := os.ReadFile("penguin.txt")
	if err != nil {
		fmt.Println(err)
	}
	return "\n" + string(penguin) + "\n"
}
