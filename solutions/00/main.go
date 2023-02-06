package main

import (
	"io"
	"log"
	"net"
)

const port = "10123"

func main() {
	srv, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()

	log.Println("listening on", port)

	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("connection from", conn.RemoteAddr())

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}
