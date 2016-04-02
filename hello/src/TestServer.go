package main

import (
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	io.Copy(conn, conn)
	conn.Close()
}

func main() {

	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}

}
