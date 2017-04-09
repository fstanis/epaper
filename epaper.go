package epaper

import "github.com/tarm/serial"

const defaultBaud = 115200

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

func (c *Client) readResponse() (string, error) {
	buf := make([]byte, 128)
	n, err := c.port.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func (c *Client) SendCommand(frame *Frame) (string, error) {
	_, err := c.port.Write(frame.Build())
	if err != nil {
		return "", err
	}

	return c.readResponse()
}
