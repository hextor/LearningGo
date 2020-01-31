package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // listens for incoming connections
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
        if err != nil {
			log.Print(err) // connection aborted
			continue
		}
		handleConn(conn) // handles a connection one at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close() // closing connection after everything run
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) // format depends on this constant `Mon Jan 2 15:04:05 MST 2006`
		if err != nil {
            return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}