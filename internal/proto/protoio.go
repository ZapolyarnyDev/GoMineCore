package proto

import (
	"errors"
	"io"
	"math"
)

const maxVarIntBytes = 5

type ProtoIO struct {
	r io.Reader
	w io.Writer
}

func New(r io.Reader, w io.Writer) *ProtoIO {
	return &ProtoIO{r: r, w: w}
}

func (p *ProtoIO) ReadVarInt() (int32, error) {
	var num int32
	var shift uint

	for i := 0; i < maxVarIntBytes; i++ {
		var b [1]byte
		if _, err := io.ReadFull(p.r, b[:]); err != nil {
			return 0, err
		}

		num |= int32(b[0]&0x7F) << shift
		if b[0]&0x80 == 0 {
			return num, nil
		}
		shift += 7
	}

	return 0, errors.New("VarInt too big")
}

func (p *ProtoIO) WriteVarInt(v int32) error {
	for {
		b := byte(v & 0x7F)
		v >>= 7
		if v != 0 {
			b |= 0x80
		}

		if _, err := p.w.Write([]byte{b}); err != nil {
			return err
		}

		if v == 0 {
			return nil
		}
	}
}

func (p *ProtoIO) ReadInt32() (int32, error) {
	var b [4]byte
	if _, err := io.ReadFull(p.r, b[:]); err != nil {
		return 0, err
	}

	return int32(b[0])<<24 |
		int32(b[1])<<16 |
		int32(b[2])<<8 |
		int32(b[3]), nil
}

func (p *ProtoIO) WriteInt32(v int32) error {
	var b [4]byte
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)

	_, err := p.w.Write(b[:])
	return err
}

func (p *ProtoIO) ReadUint32() (uint32, error) {
	var b [4]byte
	if _, err := io.ReadFull(p.r, b[:]); err != nil {
		return 0, err
	}

	return uint32(b[0])<<24 |
		uint32(b[1])<<16 |
		uint32(b[2])<<8 |
		uint32(b[3]), nil
}

func (p *ProtoIO) WriteUint32(v uint32) error {
	var b [4]byte
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)

	_, err := p.w.Write(b[:])
	return err
}

func (p *ProtoIO) ReadFloat32() (float32, error) {
	u, err := p.ReadUint32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(u), nil
}

func (p *ProtoIO) WriteFloat32(v float32) error {
	return p.WriteUint32(math.Float32bits(v))
}

func (p *ProtoIO) ReadString() (string, error) {
	length, err := p.ReadVarInt()
	if err != nil {
		return "", err
	}

	if length < 0 {
		return "", errors.New("negative string length")
	}

	buf := make([]byte, length)
	if _, err := io.ReadFull(p.r, buf); err != nil {
		return "", err
	}

	return string(buf), nil
}

func (p *ProtoIO) WriteString(s string) error {
	data := []byte(s)

	if err := p.WriteVarInt(int32(len(data))); err != nil {
		return err
	}

	_, err := p.w.Write(data)
	return err
}

func (p *ProtoIO) ReadUShort() (uint16, error) {
	var b [2]byte
	if _, err := io.ReadFull(p.r, b[:]); err != nil {
		return 0, err
	}
	return uint16(b[0])<<8 | uint16(b[1]), nil
}

func (p *ProtoIO) WriteUShort(v uint16) error {
	b := []byte{
		byte(v >> 8),
		byte(v),
	}
	_, err := p.w.Write(b)
	return err
}

func (p *ProtoIO) ReadInt64() (int64, error) {
	var b [8]byte
	if _, err := io.ReadFull(p.r, b[:]); err != nil {
		return 0, err
	}

	return int64(b[0])<<56 |
		int64(b[1])<<48 |
		int64(b[2])<<40 |
		int64(b[3])<<32 |
		int64(b[4])<<24 |
		int64(b[5])<<16 |
		int64(b[6])<<8 |
		int64(b[7]), nil
}

func (p *ProtoIO) WriteInt64(v int64) error {
	var b [8]byte
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)

	_, err := p.w.Write(b[:])
	return err
}
