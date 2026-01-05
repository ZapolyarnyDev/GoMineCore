package packet

import "gominecore/internal/proto"

type Handshake struct {
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       int32
}

func (p *Handshake) Decode(r *proto.ProtoIO) error {
	var err error

	if p.ProtocolVersion, err = r.ReadVarInt(); err != nil {
		return err
	}

	if p.ServerAddress, err = r.ReadString(); err != nil {
		return err
	}

	if p.ServerPort, err = r.ReadUShort(); err != nil {
		return err
	}

	if p.NextState, err = r.ReadVarInt(); err != nil {
		return err
	}

	return nil
}
