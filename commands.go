package epaper

func (c *Client) Handshake() error {
	frm := newFrame(commandTypeHandshake)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) SetBaudRate(rate int) error {
	frm := newFrame(commandTypeSetBaudRate, uint32(rate))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) GetBaudRate() (int, error) {
	frm := newFrame(commandTypeGetBaudRate)
	return parseInt(parseResponse(c.sendCommand(frm)))
}

func (c *Client) IsStorageSDCard() (bool, error) {
	frm := newFrame(commandTypeGetStorageType)
	return parseBool(parseResponse(c.sendCommand(frm)))
}

func (c *Client) SetStorageType(useMicroSD bool) error {
	frm := newFrame(commandTypeSetStorageType, boolToByte(useMicroSD))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) Sleep() error {
	frm := newFrame(commandTypeSleep)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) Update() error {
	frm := newFrame(commandTypeUpdate)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) IsRotated() (bool, error) {
	frm := newFrame(commandTypeGetRotation)
	return parseBool(parseResponse(c.sendCommand(frm)))
}

func (c *Client) SetRotation(rotated bool) error {
	frm := newFrame(commandTypeSetRotation, boolToByte(rotated))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) LoadFonts() error {
	frm := newFrame(commandTypeLoadFonts)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) LoadImages() error {
	frm := newFrame(commandTypeLoadImages)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) SetColor(foreground Color, background Color) error {
	frm := newFrame(commandTypeSetColor, byte(foreground), byte(background))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) GetColor() (Color, Color, error) {
	frm := newFrame(commandTypeGetColor)
	return parseColors(parseResponse(c.sendCommand(frm)))
}

func (c *Client) GetEnglishFontSize() (FontSize, error) {
	frm := newFrame(commandTypeGetEnglishFontSize)
	return parseFontSize(parseResponse(c.sendCommand(frm)))
}

func (c *Client) GetChineseFontSize() (FontSize, error) {
	frm := newFrame(commandTypeGetChineseFontSize)
	return parseFontSize(parseResponse(c.sendCommand(frm)))
}

func (c *Client) SetEnglishFontSize(size FontSize) error {
	frm := newFrame(commandTypeSetEnglishFontSize, byte(size))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) SetChineseFontSize(size FontSize) error {
	frm := newFrame(commandTypeSetChineseFontSize, byte(size))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) FillPixel(x uint16, y uint16) error {
	frm := newFrame(commandTypeFillPixel, x, y)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) DrawLine(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error {
	frm := newFrame(commandTypeDrawLine, x1, y1, x2, y2)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) FillRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error {
	frm := newFrame(commandTypeFillRect, x1, y1, x2, y2)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) DrawRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error {
	frm := newFrame(commandTypeDrawRect, x1, y1, x2, y2)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) DrawCircle(x uint16, y uint16, r uint16) error {
	frm := newFrame(commandTypeDrawCircle, x, y, r)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) FillCircle(x uint16, y uint16, r uint16) error {
	frm := newFrame(commandTypeFillCircle, x, y, r)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) DrawTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) error {
	frm := newFrame(commandTypeDrawTriangle, x1, y1, x2, y2, x3, y3)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) FillTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) error {
	frm := newFrame(commandTypeFillTriangle, x1, y1, x2, y2, x3, y3)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) Clear() error {
	frm := newFrame(commandTypeClear)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) DisplayText(x uint16, y uint16, text string) error {
	frm := newFrame(commandTypeDisplayText, x, y, text)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

func (c *Client) DisplayImage(x uint16, y uint16, file string) error {
	frm := newFrame(commandTypeDisplayImage, x, y, file)
	return wantOK(parseResponse(c.sendCommand(frm)))
}
