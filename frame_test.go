package epaper

import (
	"reflect"
	"testing"
)

func TestHandshake(t *testing.T) {
	frm := &frame{}
	frm.CommandType = commandTypeHandshake
	data := frm.Build()

	want := []byte{0xA5, 0x00, 0x09, 0x00, 0xCC, 0x33, 0xC3, 0x3C, 0xAC}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("Got %X, want %X", data, want)
	}
}

func TestLine(t *testing.T) {
	frm := &frame{}
	frm.CommandType = commandTypeDrawLine
	frm.DataAddShort(20)
	frm.DataAddShort(45)
	frm.DataAddShort(80)
	frm.DataAddShort(70)
	data := frm.Build()

	want := []byte{0xA5, 0x00, 0x11, 0x22, 0x00, 0x14, 0x00, 0x2D, 0x00, 0x50, 0x00, 0x46, 0xCC, 0x33, 0xC3, 0x3C, 0xB9}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("Got %X, want %X", data, want)
	}
}

func TestText(t *testing.T) {
	frm := &frame{}
	frm.CommandType = commandTypeDisplayText
	frm.DataAddShort(20)
	frm.DataAddShort(45)
	frm.DataAddString("test")
	data := frm.Build()

	want := []byte{0xA5, 0x00, 0x12, 0x30, 0x00, 0x14, 0x00, 0x2D, 0x74, 0x65, 0x73, 0x74, 0x00, 0xCC, 0x33, 0xC3, 0x3C, 0xA8}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("Got %X, want %X", data, want)
	}
}

func TestFrame(t *testing.T) {
	frm := newFrame(commandTypeDisplayText, uint16(20), uint16(45), "test")
	data := frm.Build()

	want := []byte{0xA5, 0x00, 0x12, 0x30, 0x00, 0x14, 0x00, 0x2D, 0x74, 0x65, 0x73, 0x74, 0x00, 0xCC, 0x33, 0xC3, 0x3C, 0xA8}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("Got %X, want %X", data, want)
	}
}
