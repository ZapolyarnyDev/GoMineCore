package session

import (
	"log"
	stdnet "net"

	"gominecore/internal/packet"
	"gominecore/internal/proto"
)

type Session struct {
	Conn  stdnet.Conn
	Proto *proto.ProtoIO
	State State
}

func New(conn stdnet.Conn) *Session {
	return &Session{
		Conn:  conn,
		Proto: proto.New(conn, conn),
		State: Handshake,
	}
}

func (s *Session) Run() {
	defer s.Conn.Close()

	for {
		frame, err := s.Proto.ReadFrame()
		if err != nil {
			log.Printf("Client %s disconnected (%v)\n", s.Conn.RemoteAddr(), err)
			return
		}

		switch s.State {

		case Handshake:
			s.handleHandshake(frame)

		case Status:
			s.handleStatus(frame)

		default:
			log.Printf("Unhandled state %v packet 0x%02X\n", s.State, frame.ID)
		}
	}
}

func (s *Session) handleHandshake(frame *proto.Frame) {
	if frame.ID != 0x00 {
		log.Printf("Unexpected packet 0x%02X in HANDSHAKE\n", frame.ID)
		return
	}

	var hs packet.Handshake
	if err := hs.Decode(frame.Body); err != nil {
		log.Printf("Handshake decode error: %v\n", err)
		return
	}

	log.Printf(
		"Handshake: proto=%d addr=%s port=%d nextState=%d\n",
		hs.ProtocolVersion,
		hs.ServerAddress,
		hs.ServerPort,
		hs.NextState,
	)

	switch hs.NextState {
	case 1:
		s.State = Status
	case 2:
		s.State = Login
	default:
		log.Printf("Invalid nextState %d\n", hs.NextState)
	}
}

func (s *Session) handleStatus(frame *proto.Frame) {
	switch frame.ID {

	case 0x00:
		resp := packet.NewStatusResponse()
		_ = s.Proto.WriteFrame(0x00, resp.Encode)

	case 0x01:
		var ping packet.Ping
		if err := ping.Decode(frame.Body); err != nil {
			return
		}
		_ = s.Proto.WriteFrame(0x01, ping.Encode)

	default:
		log.Printf("Unknown Status packet 0x%02X\n", frame.ID)
	}
}
