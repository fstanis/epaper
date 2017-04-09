package epaper

import "github.com/tarm/serial"

const (
	defaultBaud     = 115200
	maxResponseSize = 32
)

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
