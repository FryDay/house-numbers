package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func init() {
	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = brightness
	opt.Channels[0].LedCount = ledCount

	neopixel, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		panic(err)
	}
	err = neopixel.Init()
	if err != nil {
		panic(err)
	}
}

func setColor(hex string) {
	hex = strings.TrimPrefix(hex, "#")
	dec, _ := strconv.ParseUint(fmt.Sprintf("0x%s", hex), 0, 32)

	log.Printf("Color set to Hex: %s, Dec: %d\n", hex, dec)
}
