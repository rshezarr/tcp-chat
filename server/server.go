package server

import (
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
