package main

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func ListenWithTimeout() {
	listener := *createListener()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go ListenerReader(&conn)
	}
}

func ListenerReader(newConn *net.Conn) {
	conn := *newConn
	defer conn.Close()
	totalBytes := 0
	for {
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		var buf [12]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("ConnectionError %v", err)
			break
		}

		totalBytes += n
		os.Stdout.Write(buf[:n])

		token := string(buf[:n])
		token = strings.Trim(token, "\n")
		// log.Printf("Token: %q", token)
		if token == "quit" {
			break
		}
	}
	log.Printf("Received %d bytes", totalBytes)
}
