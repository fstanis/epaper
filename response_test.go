package epaper

import (
	"errors"
	"testing"
)

func TestOK(t *testing.T) {
	err := wantOK("OK", nil)
	if err != nil {
		t.Errorf("OK resulted in error: %v", err)
	}

	err = wantOK("TEST", nil)
	if err == nil {
		t.Error("not OK input produced no error")
	}

	e := errors.New("")
	err = wantOK("OK", e)
	if err != e {
		t.Error("error should have been forwarded")
	}
}

func TestError(t *testing.T) {
	testMap := map[string]error{
		"Error:0":   ErrInvalidCommand,
		"Error:4":   ErrFileNotFound,
		"Error:250": ErrUndefined,
		"Error:123": ErrUnknown,
	}

	for input, want := range testMap {
		got := parseError(input)
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}

func TestFontSize(t *testing.T) {
	s, _ := parseFontSize("1", nil)
	if s != FontSize32 {
		t.Errorf("want FontSize32, got %v", s)
	}
	_, err := parseFontSize("0", nil)
	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestColors(t *testing.T) {
	f, b, _ := parseColors("12", nil)
	if f != ColorDarkGray || b != ColorGray {
		t.Errorf("want ColorGray, ColorDarkGray, got %v, %v", f, b)
	}

	_, _, err := parseColors("12345", nil)
	if err == nil {
		t.Error("expected error for invalid input")
	}
}
