package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Color ...
type Color struct {
	Hex string `json:"hex"`
	Xeh string `json:"xeh"`
	Hue uint16 `json:"hue"`
}

// NewColor ...
func NewColor(hue uint16) *Color {
	// HSL to RGB
	var r, g, b uint8
	r = hueToRGB(float64(hue)/359.0 + 1.0/3.0)
	g = hueToRGB(float64(hue) / 359.0)
	b = hueToRGB(float64(hue)/359.0 - 1.0/3.0)

	// RGB to HEX
	hex := fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)
	xeh := fmt.Sprintf("#%.2x%.2x%.2x", 255-r, 255-g, 255-b)

	return &Color{
		Hex: hex,
		Xeh: xeh,
		Hue: hue,
	}
}

// UnmarshalJSON ...
func (color *Color) UnmarshalJSON(data []byte) error {
	hue := &struct {
		Hue string `json:"hue"`
	}{}
	if err := json.Unmarshal(data, hue); err != nil {
		return err
	}
	h, _ := strconv.Atoi(hue.Hue)
	*color = *NewColor(uint16(h))
	return nil
}

func hueToRGB(hue float64) uint8 {
	if hue < 0 {
		hue++
	}
	if hue > 1 {
		hue--
	}
	if hue < 1.0/6.0 {
		return uint8(6.0 * hue * 255.0)
	}
	if hue < 1.0/2.0 {
		return 255
	}
	if hue < 2.0/3.0 {
		return uint8((2.0/3.0 - hue) * 6.0 * 255.0)
	}
	return 0
}
