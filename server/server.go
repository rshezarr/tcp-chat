package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
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

func (h *Hub) getMessage(conn net.Conn, newConn *bufio.Reader, name string) (string, error) {
	var msg string
	var err error
	for {
		msg, err = newConn.ReadString('\n')
		userTime := time.Now().Format("2006-01-02 15:04:05")
		if err != nil {
			// If someone left chat
			return "", err
		} else if !isValidStr(msg) {
			// If message contains invalid characters
			conn.Write([]byte("[" + userTime + "] " + name + ":"))
		} else {
			// If everything is OK
			return msg, nil
		}
	}
}

func (h *Hub) sendMessage(conn net.Conn, msg string) {
	// Add all messages in txt file
	h.tempHistory = append(h.tempHistory, []byte(msg)...)
	os.WriteFile("./assets/chatHistory.txt", h.tempHistory, 0o666)
	// Send message for all online users
	h.Lock()
	for user, n := range h.users {
		if user != conn {
			user.Write([]byte(msg))
		}
		user.Write([]byte(fmt.Sprintf("[%s] [%s]: ", time.Now().Format("2006-01-02 15:04:05"), n)))
	}
	h.Unlock()
}
