package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

func (cr *ClassReader) readUint8() uint8 {
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

func (cr *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

func (cr *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

func (cr *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUint16()
	val := make([]uint16, n)
	for i, _ := range val {
		val[i] = cr.readUint16()
	}
	return val
}

func (cr *ClassReader) readBytes(n uint32) []byte {
	val := cr.data[:n]
	cr.data = cr.data[n:]
	return val
}
