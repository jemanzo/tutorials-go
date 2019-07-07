package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"
)

func CreateTLSProxy() {
	listener := *createListener()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go logSNI(conn)
	}
}

// https://youtu.be/afSiVelXDTQ
func logSNI(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, conn, 1+2+20); err != nil {
		log.Println(err)
		return
	}
	length := binary.BigEndian.Uint16(buf.Bytes()[3:5])
	if _, err := io.CopyN(&buf, conn, int64(length)); err != nil {
		log.Println(err)
		return
	}
	// ch, ok := ParseClientHello(buf.Bytes())
	// if ok {
	// 	log.Printf("SNI Connection %q", ch.SNI)
	// }
}
