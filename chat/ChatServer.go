package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println(err.Error())
	}

	// We need channels for connection, dead connection, and message
	aconns := make(map[net.Conn]int)
	conns := make(chan net.Conn)
	dconns := make(chan net.Conn)
	msgs := make(chan string)

	i := 0

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err.Error())
			}
			conns <- conn
		}
	}()

	for {
		select {
		// read incoming connection
		case conn := <-conns:

			aconns[conn] = i
			i++

			// Once we have the connection, start reading message from it
			go func(conn net.Conn, i int) {

				rd := bufio.NewReader(conn)
				for {
					m, err := rd.ReadString('\n')
					if err != nil {
						break
					}

					msgs <- fmt.Sprintf("Client %v: %v", i, m)
				}
				// Done reading from it
				dconns <- conn
			}(conn, i)

		case msg := <-msgs:
			// we have to broadcast it to all connections
			for conn := range aconns {
				_, _ = conn.Write([]byte(msg))
			}
		case dconn := <-dconns:
			log.Printf("Client %v is gone", aconns[dconn])
			delete(aconns, dconn)
		}
	}

}
