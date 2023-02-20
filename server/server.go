package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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

// Starting new server and accepting new users
func (h *Hub) Run(port string) error {
	// Check port for nonNumeric characters
	if !isValidPort(port) {
		return errors.New("invalid port")
	}
	// Starting server
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return errors.New("server start error")
	}
	defer listener.Close()
	fmt.Printf("Server is listening port :%s\n", port)
	// Accepting new users and
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection failed %v", err)
			continue
		}
		// Check users quantity. Maximum must be 10
		h.Lock()
		if len(h.users) >= maxConn {
			log.Printf("Server is full")
			conn.Write([]byte("Connection limit reached. Try again later"))
			conn.Close()
		} else {
			conn.Write(printLogo())
			conn.Write([]byte("\n[ENTER YOUR NAME]:"))
			go h.handleConn(conn)
		}
		h.Unlock()
	}
}

func (h *Hub) handleConn(conn net.Conn) {
	newConn := bufio.NewReader(conn)
	// Getting name from user and checking it for a invalid characters
	userName := strings.TrimSpace(h.getName(conn, newConn))
	log.Printf("Connection from %v as %s", conn.RemoteAddr(), userName)
	// Display message history to new users
	conn.Write(h.tempHistory)
	// Add new user to the map
	h.Lock()
	h.users[conn] = userName
	h.Unlock()
	// Send message about new user
	h.sendMessage(conn, userName+userJoinedChat)
	defer conn.Close()
	for {
		// Getting message from user and checking it for a invalid characters
		msg, err := h.getMessage(conn, newConn, userName)
		if err != nil {
			log.Printf("%s left the chat", userName)
			h.sendMessage(conn, userName+userLeftChat)
			h.Lock()
			delete(h.users, conn)
			h.Unlock()
			break
		}
		// Sending message for all online users
		userTime := time.Now().Format("2006-01-02 15:04:05")
		newMsg := fmt.Sprintf("[%s] [%s]: %s", userTime, userName, msg)
		h.sendMessage(conn, newMsg)
		defer conn.Close()
		log.Printf(userName + msg)
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
