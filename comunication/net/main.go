package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// ListenerLimitedToOneClient()
	// ListenerUnlimited()
	// ListenWithTimeout()
	CreateTCPProxy()
}

func createListener() *net.Listener {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	return &listener
}

func ListenerLimitedToOneClient() {
	listener := *createListener()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// by using io.Copy() here the loop will freeze
		// and the listener.Accept() will not run until
		// the first client disconnect

		n, err := io.Copy(os.Stdout, conn)
		log.Printf("Copied %d bytes -- Error %v", n, err)
	}
}

func ListenerUnlimited() {
	listener := *createListener()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Once we Accept() the conn we MUST create a goroutine
		// otherwise we may freeze this Accept() loop
		//   Example: a client may be able to attack the server by
		//            sending an invalid or empty Header if we try
		//            to parse Headers. Just by connecting via Telnet
		//            and not sending any data would block this loop.

		go func() {
			n, err := io.Copy(os.Stdout, conn)
			log.Printf("Copied %d bytes -- Error %v", n, err)
		}()
	}
}
