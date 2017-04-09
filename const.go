package epaper

type Color byte

const (
	ColorBlack    Color = 0x00
	ColorGray     Color = 0x01
	ColorDarkGray Color = 0x02
	ColorWhite    Color = 0x03
)

type FontSize byte

const (
	FontSize32 FontSize = 0x01
	FontSize48 FontSize = 0x02
	FontSize64 FontSize = 0x03
)

const (
	frameHeader  = 0xA5
	frameFooter0 = 0xCC
	frameFooter1 = 0x33
	frameFooter2 = 0xC3
	frameFooter3 = 0x3C
)

const (
	// System control
	commandTypeHandshake      = 0x00
	commandTypeSetBaudRate    = 0x01
	commandTypeGetBaudRate    = 0x02
	commandTypeGetStorageType = 0x06
	commandTypeSetStorageType = 0x07
	commandTypeSleep          = 0x08
	commandTypeUpdate         = 0x0A
	commandTypeGetRotation    = 0x0C
	commandTypeSetRotation    = 0x0D
	commandTypeLoadFont       = 0x0E
	commandTypeLoadImage      = 0x0F

	// Display parameter configuration
	commandTypeSetColor           = 0x10
	commandTypeGetColor           = 0x11
	commandTypeGetEnglishFontSize = 0x1C
	commandTypeGetChineseFontSize = 0x1D
	commandTypeSetEnglishFontSize = 0x1E
	commandTypeSetChineseFontSize = 0x1F

	// Basic drawings
	commandTypeFillPixel    = 0x20
	commandTypeDrawLine     = 0x22
	commandTypeFillRect     = 0x24
	commandTypeDrawRect     = 0x25
	commandTypeDrawCircle   = 0x26
	commandTypeFillCircle   = 0x27
	commandTypeDrawTriangle = 0x28
	commandTypeFillTriangle = 0x29
	commandTypeClear        = 0x2E

	// Display text
	commandTypeDisplayText = 0x30

	// Display image
	commandTypeDisplayImage = 0x70
)
