package main

import (
	"fmt"
)

// Color ...
type Color struct {
	R   uint8  `json:"r"`
	G   uint8  `json:"g"`
	B   uint8  `json:"b"`
	Hex string `json:"hex"`
}

// NewFromHex ...
func NewFromHex(hex string) *Color {
	color := &Color{
		Hex: hex,
	}
	fmt.Sscanf(hex, "#%02x%02x%02x", &color.R, &color.G, &color.B)

	return color
}

// NewFromRGB ...
func NewFromRGB(r, g, b uint8) *Color {
	return &Color{
		R:   r,
		G:   g,
		B:   b,
		Hex: fmt.Sprintf("#%.2x%.2x%.2x", r, g, b),
	}
}
