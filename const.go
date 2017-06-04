package epaper

// Color used to draw objects or as background.
type Color byte

const (
	ColorBlack     Color = 0x00
	ColorDarkGray  Color = 0x01
	ColorLightGray Color = 0x02
	ColorWhite     Color = 0x03
)

// All the colors available on the device.
var Colors = [...]Color{
	ColorBlack,
	ColorDarkGray,
	ColorLightGray,
	ColorWhite,
}

// Font size for displaying text.
type FontSize byte

const (
	FontSize32 FontSize = 0x01
	FontSize48 FontSize = 0x02
	FontSize64 FontSize = 0x03
)

// All the font sizes available on the device. `FontSize(0)` is invalid, kept
// solely to ensure array indices are aligned with the constant values.
var FontSizes = [...]FontSize{
	FontSize(0),
	FontSize32,
	FontSize48,
	FontSize64,
}

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
	commandTypeLoadFonts      = 0x0E
	commandTypeLoadImages     = 0x0F

	// Display parameter configuration
	commandTypeSetColor           = 0x10
	commandTypeGetColor           = 0x11
	commandTypeGetEnglishFontSize = 0x1C
	commandTypeGetChineseFontSize = 0x1D
	commandTypeSetEnglishFontSize = 0x1E
	commandTypeSetChineseFontSize = 0x1F

	// Basic drawings
	commandTypeFillPixel    = 0x20
	commandTypeColorPixel   = 0x21
	commandTypeDrawLine     = 0x22
	commandTypeColorLine    = 0x23
	commandTypeFillRect     = 0x24
	commandTypeDrawRect     = 0x25
	commandTypeDrawCircle   = 0x26
	commandTypeFillCircle   = 0x27
	commandTypeDrawTriangle = 0x28
	commandTypeFillTriangle = 0x29
	commandTypeClear        = 0x2E
	commandTypeFillScreen   = 0x2F

	// Display text
	commandTypeDisplayText = 0x30

	// Display image
	commandTypeDisplayImage = 0x70
)
