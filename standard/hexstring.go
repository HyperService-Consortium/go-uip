package standard

import (
	"errors"
	"unicode"
)

var (
	ErrInvalidLength           = errors.New("invalid length")
	ErrContainsInvalidHexDigit = errors.New("contains invalid hex digit")
)

func CheckValidHexString(h string) error {

	for i := range h {
		if !unicode.Is(unicode.Hex_Digit, rune(h[i])) {
			return ErrContainsInvalidHexDigit
		}
	}
	if len(h) == 0 || ((len(h)&1) != 0) {
		return ErrInvalidLength
	}
	return nil
}

func IsValidHexString(h string) bool {
	return CheckValidHexString(h) == nil
}
