package epaper

import (
	"bytes"
	"encoding/binary"
)

func parity(data []byte) byte {
	checksum := byte(0)
	for _, b := range data {
		checksum ^= b
	}
	return checksum
}

type Frame struct {
	CommandType byte
	data        bytes.Buffer
}

func (f *Frame) DataClear() {
	f.data = bytes.Buffer{}
}

func (f *Frame) DataAddByte(b byte) {
	f.data.WriteByte(b)
}

func (f *Frame) DataAddShort(short uint16) {
	binary.Write(&f.data, binary.BigEndian, short)
}

func (f *Frame) DataAddDword(dword uint32) {
	binary.Write(&f.data, binary.BigEndian, dword)
}

func (f *Frame) DataAddString(str string) {
	f.data.WriteString(str)
	f.data.WriteByte(0)
}

func (f *Frame) Length() int {
	return 9 + f.data.Len()
}

func (f *Frame) Build() []byte {
	length := f.Length()
	result := new(bytes.Buffer)
	result.Grow(length)

	result.WriteByte(frameHeader)
	binary.Write(result, binary.BigEndian, uint16(length))
	result.WriteByte(f.CommandType)
	f.data.WriteTo(result)
	result.WriteByte(frameFooter0)
	result.WriteByte(frameFooter1)
	result.WriteByte(frameFooter2)
	result.WriteByte(frameFooter3)
	result.WriteByte(parity(result.Bytes()))

	return result.Bytes()
}
