package server

import (
	"bufio"
	"log"
	"net"
	"os"
	"sync"
)

const (
	maxConn        = 10
	userJoinedChat = " has joined our chat...\n"
	userLeftChat   = " has left our chat...\n"
)

type Hub struct {
	sync.Mutex
	users       map[net.Conn]string
	tempHistory []byte
}

func NewHub() *Hub {
	return &Hub{
		users:       make(map[net.Conn]string),
		tempHistory: make([]byte, 0),
	}
}

func printLogo() []byte {
	wlcm, err := os.ReadFile("./assets/logo.txt")
	if err != nil {
		log.Printf("logo error: %v", err)
	}
	return wlcm
}

func (h *Hub) getName(conn net.Conn, newConn *bufio.Reader) string {
	for {
		name, _ := newConn.ReadString('\n')
		if !isValidStr(name) || h.isTakenName(conn, name) {
			// If username is incorrect or already exist
			conn.Write([]byte("\n[ENTER YOUR NAME CORRECTLY]:"))
		} else {
			// If everything is OK
			return name
		}
	}
}
