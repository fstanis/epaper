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
	commandType byte
	data        bytes.Buffer
}

func newFrame(commandType byte, datas ...interface{}) *frame {
	frame := &frame{commandType: commandType}
	for _, data := range datas {
		switch val := data.(type) {
		case byte:
			frame.dataAddByte(val)
		case uint16:
			frame.dataAddShort(val)
		case uint32:
			frame.dataAddDword(val)
		case string:
			frame.dataAddString(val)
		default:
			log.Fatalf("Unknown data type in frame: %T", data)
		}
	}
	return frame
}

func (f *frame) dataClear() {
	f.data = bytes.Buffer{}
}

func (f *frame) dataAddByte(b byte) {
	f.data.WriteByte(b)
}

func (f *frame) dataAddShort(short uint16) {
	binary.Write(&f.data, binary.BigEndian, short)
}

func (f *frame) dataAddDword(dword uint32) {
	binary.Write(&f.data, binary.BigEndian, dword)
}

func (f *frame) dataAddString(str string) {
	f.data.WriteString(str)
	f.data.WriteByte(0)
}

func (f *frame) length() int {
	return 9 + f.data.Len()
}

func (f *frame) build() []byte {
	length := f.length()
	result := new(bytes.Buffer)
	result.Grow(length)

	result.WriteByte(frameHeader)
	binary.Write(result, binary.BigEndian, uint16(length))
	result.WriteByte(f.commandType)
	f.data.WriteTo(result)
	result.WriteByte(frameFooter0)
	result.WriteByte(frameFooter1)
	result.WriteByte(frameFooter2)
	result.WriteByte(frameFooter3)
	result.WriteByte(parity(result.Bytes()))

	return result.Bytes()
}
