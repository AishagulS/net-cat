package net_cat

import (
	"net"
	"sync"
	"time"
)

const maxClients = 10

type client struct {
	conn net.Conn
	name string
}

type chatRoom struct {
	clients []client
	mutex   sync.Mutex
	history []string
}

var timestamp = time.Now().Format("2006-01-02 15:04:05")
