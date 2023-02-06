package main

import (
	"bufio"
	"log"
	"net"
	"solutions/01/handler"
	"solutions/01/prime"
	"solutions/01/protocol"
)

const port = "10123"

func main() {
	srv, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on", port)

	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(client net.Conn) {
			scanner := bufio.NewScanner(conn)

			for scanner.Scan() {
				line := scanner.Text()
				log.Println("line received", line)

				var res []byte
				num, err := handler.GetNumber(line)
				if err == nil {
					res = protocol.MakeResponse(prime.IsPrime(num))
				} else {
					res = protocol.MakeError(err.Error())
				}

				log.Println("response", string(res))
				client.Write(res)
				client.Write([]byte("\n"))

				if err != nil {
					break
				}
			}

			err := scanner.Err()
			if err != nil {
				log.Println("socket error", err)
			}

			client.Close()
		}(conn)
	}
}
