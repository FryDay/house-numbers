package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func setColor(hex string) {
	hex = strings.TrimPrefix(hex, "#")
	dec, _ := strconv.ParseUint(fmt.Sprintf("0x%s", hex), 0, 32)

	log.Printf("Color set to Hex: %s, Dec: %d\n", hex, dec)
}
