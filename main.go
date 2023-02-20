package main

import (
	"fmt"
	"log"
	"os"
	"tcp-chat/server"
)

func main() {
	port := "8989"
	args := os.Args[1:]
	if len(args) == 1 {
		port = args[0]
	} else if len(args) != 0 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	if err := server.NewHub().Run(port); err != nil {
		log.Printf("Error: %v", err)
	}
}
