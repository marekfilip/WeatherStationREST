package Brightness

import "strconv"

type Brightness uint16

func NewBrightness(x uint16) *Brightness {
	convertedBrightness := Brightness(x)
	return &convertedBrightness
}

func ParseString(b string) (*Brightness, error) {
	converted, err := strconv.ParseUint(b, 10, 16)
	if err != nil {
		return nil, err
	}

	return NewBrightness(uint16(converted)), nil
}
