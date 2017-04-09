package epaper

import (
	"reflect"
	"testing"
)

func TestHandshake(t *testing.T) {
	frame := &Frame{}
	frame.CommandType = CMD_HANDSHAKE
	data := frame.Build()

	want := []byte{0xA5, 0x00, 0x09, 0x00, 0xCC, 0x33, 0xC3, 0x3C, 0xAC}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("Got %X, want %X", data, want)
	}
}

func TestLine(t *testing.T) {
	frame := &Frame{}
	frame.CommandType = CMD_DRAW_LINE
	frame.DataAddShort(20)
	frame.DataAddShort(45)
	frame.DataAddShort(80)
	frame.DataAddShort(70)
	data := frame.Build()

	want := []byte{0xA5, 0x00, 0x11, 0x22, 0x00, 0x14, 0x00, 0x2D, 0x00, 0x50, 0x00, 0x46, 0xCC, 0x33, 0xC3, 0x3C, 0xB9}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("Got %X, want %X", data, want)
	}
}

func TestText(t *testing.T) {
	frame := &Frame{}
	frame.CommandType = CMD_DRAW_STRING
	frame.DataAddShort(20)
	frame.DataAddShort(45)
	frame.DataAddString("test")
	data := frame.Build()

	want := []byte{0xA5, 0x00, 0x12, 0x30, 0x00, 0x14, 0x00, 0x2D, 0x74, 0x65, 0x73, 0x74, 0x00, 0xCC, 0x33, 0xC3, 0x3C, 0xA8}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("Got %X, want %X", data, want)
	}
}
