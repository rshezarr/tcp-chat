package server

import (
	"net"
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
