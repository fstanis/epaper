package epaper

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	responseOK = "OK"
)

var (
	ErrInvalidCommand     = errors.New("invalid command (error 0)")
	ErrSDInitiationFailed = errors.New("SD card initiation failed (error 1)")
	ErrInvalidArguments   = errors.New("invalid arguments (error 2)")
	ErrNoSD               = errors.New("SD card not inserted (error 3)")
	ErrFileNotFound       = errors.New("file not found (error 4)")
	ErrValidationFailed   = errors.New("validation failed (error 20)")
	ErrInvalidFrame       = errors.New("invalid frame (error 21)")
	ErrUndefined          = errors.New("undefined error (error 250)")
	ErrUnknown            = errors.New("unknown error")
)

var (
	regexError = regexp.MustCompile(`^Error:(\d+)$`)

	commandErrors = map[int]error{
		0:   ErrInvalidCommand,
		1:   ErrSDInitiationFailed,
		2:   ErrInvalidArguments,
		3:   ErrNoSD,
		4:   ErrFileNotFound,
		20:  ErrValidationFailed,
		21:  ErrInvalidFrame,
		250: ErrUndefined,
	}

	colors = map[byte]Color{
		'0': ColorBlack,
		'1': ColorDarkGray,
		'2': ColorGray,
		'3': ColorWhite,
	}

	fontSizes = map[string]FontSize{
		"1": FontSize32,
		"2": FontSize48,
		"3": FontSize64,
	}
)

func parseError(resp string) error {
	matches := regexError.FindStringSubmatch(resp)
	if len(matches) == 2 {
		errID, _ := strconv.Atoi(matches[1])
		err, hasErr := commandErrors[errID]
		if !hasErr {
			err = ErrUnknown
		}
		return err
	}
	return nil
}

func parseResponse(resp string, err error) (string, error) {
	if err != nil {
		return "", fmt.Errorf("communication error: %v")
	}

	if err = parseError(resp); err != nil {
		return "", err
	}
	return resp, nil
}

func wantOK(resp string, err error) error {
	if err != nil {
		return err
	}

	if resp != responseOK {
		return fmt.Errorf("command failed: returned %q, want %q", resp, responseOK)
	}
	return nil
}

func parseInt(resp string, err error) (int, error) {
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(resp)
	if err != nil {
		return 0, fmt.Errorf("invalid response, expected number, got: %q", resp)
	}
	return i, nil
}

func parseBool(resp string, err error) (bool, error) {
	if err != nil {
		return false, err
	}

	switch resp {
	case "0":
		return false, nil
	case "1":
		return true, nil
	default:
		return false, fmt.Errorf("invalid response, expected 0 or 1, got: %q", resp)
	}
}

func parseFontSize(resp string, err error) (FontSize, error) {
	if err != nil {
		return 0, err
	}

	size, ok := fontSizes[resp]
	if !ok {
		return 0, fmt.Errorf("invalid response, got %q", resp)
	}
	return FontSize(size), nil
}

func parseColors(resp string, err error) (Color, Color, error) {
	if err != nil {
		return 0, 0, err
	}

	if len(resp) != 2 {
		return 0, 0, fmt.Errorf("invalid response, got %q", resp)
	}
	f, okf := colors[resp[0]]
	b, okb := colors[resp[1]]
	if !okf || !okb {
		return 0, 0, fmt.Errorf("invalid response, got %q", resp)
	}

	return f, b, nil
}

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}
