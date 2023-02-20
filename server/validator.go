package server

import (
	"net"
	"strings"
)

// Checking messages and names function
func isValidStr(str string) bool {
	for _, w := range str {
		if w != ' ' && w != '\n' {
			return true
		}
	}
	return false
}

// Checking port function
func isValidPort(port string) bool {
	for _, v := range port {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

// Check for taken names
func (h *Hub) isTakenName(conn net.Conn, tempName string) bool {
	h.Lock()
	for _, u := range h.users {
		if u == strings.TrimSpace(tempName) {
			conn.Write([]byte("Name has already taken! Try again"))
			return true
		}
	}
	h.Unlock()
	return false
}
