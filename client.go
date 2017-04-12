package epaper

import "github.com/tarm/serial"

const (
	defaultBaud     = 115200
	maxResponseSize = 32
)

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
	DrawLine(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error
	FillRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error
	DrawRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error
	DrawCircle(x uint16, y uint16, r uint16) error
	FillCircle(x uint16, y uint16, r uint16) error
	DrawTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) error
	FillTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) error
	Clear() error
	DisplayText(x uint16, y uint16, text string) error
	DisplayImage(x uint16, y uint16, file string) error
	Close() error
}

type Client struct {
	port *serial.Port
}

func New(port string) (*Client, error) {
	return NewWithBaud(port, defaultBaud)
}

func NewWithBaud(port string, baud int) (*Client, error) {
	s, err := serial.OpenPort(&serial.Config{Name: port, Baud: baud})
	if err != nil {
		return nil, err
	}
	return &Client{s}, nil
}

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
	_, err := c.port.Write(frm.Build())
	if err != nil {
		return "", err
	}

	return c.readResponse()
}
