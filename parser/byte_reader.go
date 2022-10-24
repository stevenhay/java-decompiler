package parser

import (
	"bytes"
	"encoding/binary"
)

type ByteReader struct {
	bytes *bytes.Buffer
}

func NewByteReader(b []byte) *ByteReader {
	return &ByteReader{
		bytes: bytes.NewBuffer(b),
	}
}

func (r *ByteReader) ReadUint8() uint8 {
	return r.Read(1)[0]
}

func (r *ByteReader) ReadUint16() uint16 {
	return binary.BigEndian.Uint16(r.Read(2))
}

func (r *ByteReader) ReadUint32() uint32 {
	return binary.BigEndian.Uint32(r.Read(4))
}

func (r *ByteReader) Read(p int) []byte {
	b := make([]byte, p)
	r.bytes.Read(b)
	return b
}

func (r *ByteReader) Len() int {
	return r.bytes.Len()
}
