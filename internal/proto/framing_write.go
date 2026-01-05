package proto

import "bytes"

func (p *ProtoIO) WriteFrame(packetId int32, encode func(*ProtoIO) error) error {
	var buf bytes.Buffer
	body := New(nil, &buf)

	if err := body.WriteVarInt(packetId); err != nil {
		return err
	}

	if err := encode(body); err != nil {
		return err
	}

	if err := p.WriteVarInt(int32(buf.Len())); err != nil {
		return err
	}

	_, err := p.w.Write(buf.Bytes())
	return err
}
