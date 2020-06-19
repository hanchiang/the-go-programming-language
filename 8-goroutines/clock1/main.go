// A TCP server that periodically writes the time

package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// blocks until an incoming connection request is made
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g. connection aborted
			continue
		}
		// handleConn(conn) // handle one connection at a time
		go handleConn(conn) // handle multiple connections at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// Since net.Conn satisfies the io.Writer interface, we can
		// write directly to it.
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g. client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
