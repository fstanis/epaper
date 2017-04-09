package epaper

import (
	"bytes"
	"encoding/binary"
	"log"
)

func parity(data []byte) byte {
	checksum := byte(0)
	for _, b := range data {
		checksum ^= b
	}
	return checksum
}

type frame struct {
	CommandType byte
	data        bytes.Buffer
}

func newFrame(commandType byte, datas ...interface{}) *frame {
	frame := &frame{CommandType: commandType}
	for _, data := range datas {
		switch val := data.(type) {
		case byte:
			frame.DataAddByte(val)
		case uint16:
			frame.DataAddShort(val)
		case uint32:
			frame.DataAddDword(val)
		case string:
			frame.DataAddString(val)
		default:
			log.Fatalf("Unknown data type in frame: %T", data)
		}
	}
	return frame
}

func (f *frame) DataClear() {
	f.data = bytes.Buffer{}
}

func (f *frame) DataAddByte(b byte) {
	f.data.WriteByte(b)
}

func (f *frame) DataAddShort(short uint16) {
	binary.Write(&f.data, binary.BigEndian, short)
}

func (f *frame) DataAddDword(dword uint32) {
	binary.Write(&f.data, binary.BigEndian, dword)
}

func (f *frame) DataAddString(str string) {
	f.data.WriteString(str)
	f.data.WriteByte(0)
}

func (f *frame) Length() int {
	return 9 + f.data.Len()
}

func (f *frame) Build() []byte {
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
