package main

import (
	"fmt"
	"strconv"
	"strings"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func init() {
	var err error

	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = brightness
	opt.Channels[0].LedCount = ledCount

	neopixel, err = ws2811.MakeWS2811(&opt)
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

	for led := 0; led < ledCount; led++ {
		neopixel.Leds(0)[led] = uint32(dec)
	}

	if err := neopixel.Render(); err != nil {
		// TODO: Handle this gracefully
		panic(err)
	}
}
