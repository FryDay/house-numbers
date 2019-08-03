package main

import (
	"encoding/json"
	"fmt"
)

// Color ...
type Color struct {
	R   uint8  `json:"r"`
	G   uint8  `json:"g"`
	B   uint8  `json:"b"`
	Hex string `json:"hex"`
}

// NewColorFromHex ...
func NewColorFromHex(hex string) *Color {
	color := &Color{
		Hex: hex,
	}
	fmt.Sscanf(hex, "#%02x%02x%02x", &color.R, &color.G, &color.B)

	return color
}

// NewColorFromRGB ...
func NewColorFromRGB(r, g, b uint8) *Color {
	return &Color{
		R:   r,
		G:   g,
		B:   b,
		Hex: fmt.Sprintf("#%.2x%.2x%.2x", r, g, b),
	}
}

// UnmarshalJSON ...
func (color *Color) UnmarshalJSON(data []byte) error {
	hex := &struct {
		Hex string `json:"hex"`
	}{}
	if err := json.Unmarshal(data, hex); err != nil {
		return err
	}
	*color = *NewColorFromHex(hex.Hex)
	return nil
}
