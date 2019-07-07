package main

import (
	"io"
	"log"
	"net"
)

func CreateTCPProxy() {
	listener := *createListener()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	defer conn.Close()
	remote, err := net.Dial("tcp", "www.goggle.com:80")
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()
	go io.Copy(remote, conn)
	io.Copy(conn, remote)
}
