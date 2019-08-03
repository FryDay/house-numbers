package main

import (
	"encoding/json"
	"fmt"
)

// Color ...
type Color struct {
	Hex string `json:"hex"`
	Xeh string `json:"xeh"`
}

// NewColor ...
func NewColor(hex string) *Color {
	color := &Color{
		Hex: hex,
	}
	var r, g, b uint8
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	color.Xeh = fmt.Sprintf("#%.2x%.2x%.2x", 255-r, 255-g, 255-b)

	return color
}

// UnmarshalJSON ...
func (color *Color) UnmarshalJSON(data []byte) error {
	hex := &struct {
		Hex string `json:"hex"`
	}{}
	if err := json.Unmarshal(data, hex); err != nil {
		return err
	}
	*color = *NewColor(hex.Hex)
	return nil
}
