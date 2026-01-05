package proto

import (
	"errors"
	"io"
)

type Frame struct {
	Length int32
	ID     int32
	Body   *ProtoIO
}

func (p *ProtoIO) ReadFrame() (*Frame, error) {
	length, err := p.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if length <= 0 {
		return nil, errors.New("invalid packet length")
	}

	lr := &io.LimitedReader{R: p.r, N: int64(length)}

	in := New(lr, nil)
	id, err := in.ReadVarInt()
	if err != nil {
		return nil, err
	}

	body := New(lr, nil)

	return &Frame{
		Length: length,
		ID:     id,
		Body:   body,
	}, nil
}
