package main

import (
	"log"
	stdnet "net"

	netlayer "gominecore/internal/net"
	"gominecore/internal/session"
)

func main() {
	log.Println("GoMineCore v0.1.0")
	log.Println("Minecraft Server starting...")

	addr := netlayer.Address{
		IP:   "0.0.0.0",
		Port: 37029,
	}

	err := netlayer.Listen(addr, func(conn stdnet.Conn) {
		s := session.New(conn)
		s.Run()
	})

	if err != nil {
		log.Fatal(err)
	}
}
