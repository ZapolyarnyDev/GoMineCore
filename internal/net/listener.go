package net

import (
	"log"
	"net"
	"strconv"
)

type Handler func(conn net.Conn)
type Address struct {
	IP   string
	Port int
}

func (a Address) String() string {
	return a.IP + ":" + strconv.Itoa(a.Port)
}

func Listen(addr Address, handler Handler) error {
	listener, err := net.Listen("tcp", addr.String())

	if err != nil {
		return err
	}

	log.Printf("Listening on %s\n", addr.String())

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Listener accept error: ", err)
			return err
		}

		log.Printf("Client connected: %s\n", conn.RemoteAddr())
		go handler(conn)
	}
}
