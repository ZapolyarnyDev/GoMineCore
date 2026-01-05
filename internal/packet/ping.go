package packet

import "gominecore/internal/proto"

type Ping struct {
	Payload int64
}

func (p *Ping) Decode(r *proto.ProtoIO) error {
	var err error
	p.Payload, err = r.ReadInt64()
	return err
}

func (p *Ping) Encode(w *proto.ProtoIO) error {
	return w.WriteInt64(p.Payload)
}
