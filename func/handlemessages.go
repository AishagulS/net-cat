package net_cat

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

func handleMessages(room *chatRoom, clientIdx int) {
	client := room.clients[clientIdx]
	conn := client.conn
	reader := bufio.NewReader(conn)
	for {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(conn, "\n\033[1A"+"\033[K[%s][%s]:", timestamp, client.name+"")
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		msg = strings.TrimSpace(msg)
		if len(msg) > 0 && MsgCheck(msg) {
			room.mutex.Lock()
			room.history = append(room.history, fmt.Sprintf("[%s][%s]: %s\n", timestamp, client.name, msg))
			room.mutex.Unlock()
			for i, c := range room.clients {
				if i != clientIdx {
					c.conn.Write([]byte(fmt.Sprintf("\n\033[1A"+"\033[K[%s][%s]:%s", timestamp, client.name, msg+"\n"+"["+timestamp+"]["+client.name+"]:")))
				}
			}

		}
	}
	conn.Close()
	room.mutex.Lock()
	room.clients = append(room.clients[:clientIdx], room.clients[clientIdx+1:]...)
	room.mutex.Unlock()
	room.mutex.Lock()
	for _, c := range room.clients {
		c.conn.Write([]byte(fmt.Sprintf("\n\033[1A"+"\033[K[%s][SERVER]: %s has left the chat", timestamp, client.name+"\n"+"["+timestamp+"]["+client.name+"]:")))
	}
	room.mutex.Unlock()
}
