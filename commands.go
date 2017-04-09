package epaper

var (
	CommandHandshake      = &Frame{CommandType: commandTypeHandshake}
	CommandGetBaudRate    = &Frame{CommandType: commandTypeGetBaudRate}
	CommandGetStorageType = &Frame{CommandType: commandTypeGetStorageType}
	CommandSleep          = &Frame{CommandType: commandTypeSleep}
	CommandUpdate         = &Frame{CommandType: commandTypeUpdate}
	CommandGetRotation    = &Frame{CommandType: commandTypeGetRotation}
	CommandLoadFont       = &Frame{CommandType: commandTypeLoadFont}
	CommandLoadImage      = &Frame{CommandType: commandTypeLoadImage}

	CommandGetColor           = &Frame{CommandType: commandTypeGetColor}
	CommandGetEnglishFontSize = &Frame{CommandType: commandTypeGetEnglishFontSize}
	CommandGetChineseFontSize = &Frame{CommandType: commandTypeGetChineseFontSize}

	CommandClear = &Frame{CommandType: commandTypeClear}
)

func CommandSetBaudRate(rate uint32) *Frame {
	frame := &Frame{CommandType: commandTypeSetBaudRate}
	frame.DataAddDword(rate)
	return frame
}

func CommandSetStorageType(useMicroSD bool) *Frame {
	area := byte(0)
	if useMicroSD {
		area = 1
	}

	frame := &Frame{CommandType: commandTypeSetStorageType}
	frame.DataAddByte(area)
	return frame
}

func CommandSetRotation(rotate bool) *Frame {
	rotation := byte(0)
	if rotate {
		rotation = 1
	}

	frame := &Frame{CommandType: commandTypeSetRotation}
	frame.DataAddByte(rotation)
	return frame
}

func CommandSetColor(foreground Color, background Color) *Frame {
	frame := &Frame{CommandType: commandTypeSetColor}
	frame.DataAddByte(byte(foreground))
	frame.DataAddByte(byte(background))
	return frame
}

func CommandSetEnglishFontSize(size FontSize) *Frame {
	frame := &Frame{CommandType: commandTypeSetEnglishFontSize}
	frame.DataAddByte(byte(size))
	return frame
}

func CommandSetChineseFontSize(size FontSize) *Frame {
	frame := &Frame{CommandType: commandTypeSetChineseFontSize}
	frame.DataAddByte(byte(size))
	return frame
}

func CommandFillPixel(x uint16, y uint16) *Frame {
	frame := &Frame{CommandType: commandTypeFillPixel}
	frame.DataAddShort(x)
	frame.DataAddShort(y)
	return frame
}

func CommandDrawLine(x1 uint16, y1 uint16, x2 uint16, y2 uint16) *Frame {
	frame := &Frame{CommandType: commandTypeDrawLine}
	frame.DataAddShort(x1)
	frame.DataAddShort(y1)
	frame.DataAddShort(x2)
	frame.DataAddShort(y2)
	return frame
}

func CommandFillRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) *Frame {
	frame := &Frame{CommandType: commandTypeFillRect}
	frame.DataAddShort(x1)
	frame.DataAddShort(y1)
	frame.DataAddShort(x2)
	frame.DataAddShort(y2)
	return frame
}

func CommandDrawRect(x1 uint16, y1 uint16, x2 uint16, y2 uint16) *Frame {
	frame := &Frame{CommandType: commandTypeDrawRect}
	frame.DataAddShort(x1)
	frame.DataAddShort(y1)
	frame.DataAddShort(x2)
	frame.DataAddShort(y2)
	return frame
}

func CommandDrawCircle(x uint16, y uint16, r uint16) *Frame {
	frame := &Frame{CommandType: commandTypeDrawCircle}
	frame.DataAddShort(x)
	frame.DataAddShort(y)
	frame.DataAddShort(r)
	return frame
}

func CommandFillCircle(x uint16, y uint16, r uint16) *Frame {
	frame := &Frame{CommandType: commandTypeFillCircle}
	frame.DataAddShort(x)
	frame.DataAddShort(y)
	frame.DataAddShort(r)
	return frame
}

func CommandDrawTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) *Frame {
	frame := &Frame{CommandType: commandTypeDrawTriangle}
	frame.DataAddShort(x1)
	frame.DataAddShort(y1)
	frame.DataAddShort(x2)
	frame.DataAddShort(y2)
	frame.DataAddShort(x3)
	frame.DataAddShort(y3)
	return frame
}

func CommandFillTriangle(x1 uint16, y1 uint16, x2 uint16, y2 uint16, x3 uint16, y3 uint16) *Frame {
	frame := &Frame{CommandType: commandTypeFillTriangle}
	frame.DataAddShort(x1)
	frame.DataAddShort(y1)
	frame.DataAddShort(x2)
	frame.DataAddShort(y2)
	frame.DataAddShort(x3)
	frame.DataAddShort(y3)
	return frame
}

func CommandDisplayText(x uint16, y uint16, text string) *Frame {
	frame := &Frame{CommandType: commandTypeDisplayText}
	frame.DataAddShort(x)
	frame.DataAddShort(y)
	frame.DataAddString(text)
	return frame
}

func CommandDisplayImage(x uint16, y uint16, file string) *Frame {
	frame := &Frame{CommandType: commandTypeDisplayImage}
	frame.DataAddShort(x)
	frame.DataAddShort(y)
	frame.DataAddString(file)
	return frame
}
