/*
This is a Go library that exposes the interface of Waveshare's 4.3inch e-Paper
UART Module

Most of the interface follows the wiki page (http://www.waveshare.com/wiki/4.3inch_e-Paper),
but some of it was derived by experimenting with the device.

Please see the project page for more information on how to get started: https://github.com/fstanis/epaper
*/
package epaper

import "github.com/tarm/serial"

const (
	defaultBaud     = 115200
	maxResponseSize = 32
)

// Device is an interface representing an e-Paper device. It's implemented by
// epaper.Client.
type Device interface {
	Handshake() error
	SetBaudRate(rate int) error
	GetBaudRate() (int, error)
	IsStorageSDCard() (bool, error)
	SetStorageType(useMicroSD bool) error
	Sleep() error
	Update() error
	IsRotated() (bool, error)
	SetRotation(rotated bool) error
	LoadFonts() error
	LoadImages() error
	SetColor(foreground Color, background Color) error
	GetColor() (Color, Color, error)
	GetEnglishFontSize() (FontSize, error)
	GetChineseFontSize() (FontSize, error)
	SetEnglishFontSize(size FontSize) error
	SetChineseFontSize(size FontSize) error
	FillPixel(x uint16, y uint16) error
	ColorPixel(x uint16, y uint16, color Color) error
	DrawLine(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error
	ColorLine(x1 uint16, y1 uint16, x2 uint16, y2 uint16, color Color) error
	FillRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error
	DrawRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error
	DrawCircle(x uint16, y uint16, r uint16) error
	FillCircle(x uint16, y uint16, r uint16) error
	DrawTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) error
	FillTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) error
	Clear() error
	FillScreen(color Color) error
	DisplayText(x uint16, y uint16, text string) error
	DisplayImage(x uint16, y uint16, file string) error
	Close() error
}

// Client is used to communicate with an e-Paper device and send commands to it.
type Client struct {
	port *serial.Port
}

// New connects to the e-Paper device on the given port. It uses the default
// bitrate of 115200.
func New(port string) (*Client, error) {
	return NewWithBaud(port, defaultBaud)
}

// NewWithBaud connects to the e-Paper device on the given port using the given
// bitrate (baud rate).
func NewWithBaud(port string, baud int) (*Client, error) {
	s, err := serial.OpenPort(&serial.Config{Name: port, Baud: baud})
	if err != nil {
		return nil, err
	}
	return &Client{s}, nil
}

// Close closes the connection on this Client's port.
func (c *Client) Close() error {
	return c.port.Close()
}

func (c *Client) readResponse() (string, error) {
	buf := make([]byte, maxResponseSize)
	n, err := c.port.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func (c *Client) sendCommand(frm *frame) (string, error) {
	_, err := c.port.Write(frm.build())
	if err != nil {
		return "", err
	}

	return c.readResponse()
}
