package epaper

// Handshake command. Essentially no-op, used to confirm a connection with the
// device is established.
func (c *Client) Handshake() error {
	frm := newFrame(commandTypeHandshake)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// Sets the serial baud rate on the device. Avoid or use with caution.
func (c *Client) SetBaudRate(rate int) error {
	frm := newFrame(commandTypeSetBaudRate, uint32(rate))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// Gets the current serial baud rate.
func (c *Client) GetBaudRate() (int, error) {
	frm := newFrame(commandTypeGetBaudRate)
	return parseInt(parseResponse(c.sendCommand(frm)))
}

// IsStorageSDCard returns true if the storage type is set to SD card. The
// storage type is only relevant for the `DisplayImage` method. By default, the
// SD card is not used.
func (c *Client) IsStorageSDCard() (bool, error) {
	frm := newFrame(commandTypeGetStorageType)
	return parseBool(parseResponse(c.sendCommand(frm)))
}

// SetStorageType sets the default storage type. The storage type is only
// relevant for the `DisplayImage` method. By default, the SD card is not used.
func (c *Client) SetStorageType(useMicroSD bool) error {
	frm := newFrame(commandTypeSetStorageType, boolToByte(useMicroSD))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// Sleep makes the device enter sleep mode which disables all commands, turns
// off the LED indicator and reduces power consumption. The only way to wake up
// from sleep mode is via the WAKE_UP pin.
func (c *Client) Sleep() error {
	frm := newFrame(commandTypeSleep)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// Update makes the device the draw the shapes, images and text sent via other
// commands.
func (c *Client) Update() error {
	frm := newFrame(commandTypeUpdate)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// IsRotated returns true if the screen is currently rotated to portrait mode.
func (c *Client) IsRotated() (bool, error) {
	frm := newFrame(commandTypeGetRotation)
	return parseBool(parseResponse(c.sendCommand(frm)))
}

// SetRotation enables when the parameter is true) or disables portrait mode.
func (c *Client) SetRotation(rotated bool) error {
	frm := newFrame(commandTypeSetRotation, boolToByte(rotated))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// LoadFonts clears the fonts on the internal flash storage and copies the fonts
// from the SD card, if there are any present. Fails if SD card is not present.
func (c *Client) LoadFonts() error {
	frm := newFrame(commandTypeLoadFonts)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// LoadFonts clears the images on the internal flash storage and copies the
// images from the SD card, if there are any present. Fails if SD card is not
// present.
func (c *Client) LoadImages() error {
	frm := newFrame(commandTypeLoadImages)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// SetColor sets the foreground and background color for future commands. The
// background color is used solely by the `Clear` method while the foreground
// color is used by all other drawing methods.
func (c *Client) SetColor(foreground Color, background Color) error {
	frm := newFrame(commandTypeSetColor, byte(foreground), byte(background))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// GetColor returns both the currently set foreground and background colors.
func (c *Client) GetColor() (Color, Color, error) {
	frm := newFrame(commandTypeGetColor)
	return parseColors(parseResponse(c.sendCommand(frm)))
}

// GetEnglishFontSize returns the currently set size of the font used for
// displaying English (ASCII) characters.
func (c *Client) GetEnglishFontSize() (FontSize, error) {
	frm := newFrame(commandTypeGetEnglishFontSize)
	return parseFontSize(parseResponse(c.sendCommand(frm)))
}

// GetChineseFontSize returns the currently set size of the font used for
// displaying Chinese characters.
func (c *Client) GetChineseFontSize() (FontSize, error) {
	frm := newFrame(commandTypeGetChineseFontSize)
	return parseFontSize(parseResponse(c.sendCommand(frm)))
}

// SetEnglishFontSize sets the size of the font used for displaying English
// (ASCII) characters.
func (c *Client) SetEnglishFontSize(size FontSize) error {
	frm := newFrame(commandTypeSetEnglishFontSize, byte(size))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// SetChineseFontSize sets the size of the font used for displaying Chinese
// characters.
func (c *Client) SetChineseFontSize(size FontSize) error {
	frm := newFrame(commandTypeSetChineseFontSize, byte(size))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// FillPixel fills the pixel at the given coordinates with the foreground color.
func (c *Client) FillPixel(x uint16, y uint16) error {
	frm := newFrame(commandTypeFillPixel, x, y)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// ColorPixel fills the pixel at the given coordinates with the given color.
func (c *Client) ColorPixel(x uint16, y uint16, color Color) error {
	frm := newFrame(commandTypeColorPixel, x, y, byte(color))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// DrawLine draws a line connecting the given two points using the foreground
// color.
func (c *Client) DrawLine(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error {
	frm := newFrame(commandTypeDrawLine, x1, y1, x2, y2)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// ColorLine draws a line connecting the given two points using the given color.
func (c *Client) ColorLine(x1 uint16, y1 uint16, x2 uint16, y2 uint16, color Color) error {
	frm := newFrame(commandTypeColorLine, x1, y1, x2, y2, byte(color))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// FillRect fills a rectangle with the given top left and bottom right corner
// using the foreground color.
func (c *Client) FillRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error {
	frm := newFrame(commandTypeFillRect, x1, y1, x2, y2)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// DrawRect draws a rectangle with the given top left and bottom right corner
// using the foreground color.
func (c *Client) DrawRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) error {
	frm := newFrame(commandTypeDrawRect, x1, y1, x2, y2)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// DrawCircle draws a circle with the given center and radius using the
// foreground color.
func (c *Client) DrawCircle(x uint16, y uint16, r uint16) error {
	frm := newFrame(commandTypeDrawCircle, x, y, r)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// FillCircle fills a circle with the given center and radius using the
// foreground color.
func (c *Client) FillCircle(x uint16, y uint16, r uint16) error {
	frm := newFrame(commandTypeFillCircle, x, y, r)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// DrawTriangle draws a triangle with the given three edges using the foreground
// color.
func (c *Client) DrawTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) error {
	frm := newFrame(commandTypeDrawTriangle, x1, y1, x2, y2, x3, y3)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// FillTriangle fills a triangle with the given three edges using the foreground
// color.
func (c *Client) FillTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) error {
	frm := newFrame(commandTypeFillTriangle, x1, y1, x2, y2, x3, y3)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// Clears the screen by filling it with the background color.
func (c *Client) Clear() error {
	frm := newFrame(commandTypeClear)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// Fills the screen with the given color.
func (c *Client) FillScreen(color Color) error {
	frm := newFrame(commandTypeFillScreen, byte(color))
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// Displays the given text at the given coordinates. The coordinates will be
// used as the top left corner of the box containing the text.
func (c *Client) DisplayText(x uint16, y uint16, text string) error {
	frm := newFrame(commandTypeDisplayText, x, y, text)
	return wantOK(parseResponse(c.sendCommand(frm)))
}

// Displays the image with the given file name, loaded either from the internal
// flash memory or the SD card, depending on the selected storage type.
func (c *Client) DisplayImage(x uint16, y uint16, file string) error {
	frm := newFrame(commandTypeDisplayImage, x, y, file)
	return wantOK(parseResponse(c.sendCommand(frm)))
}
