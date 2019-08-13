package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	ledCount   = 7
	brightness = 255
	lat        = 44.280339
	lng        = -88.4018744
	loc        = "America/Chicago"
)

var (
	ledsOn       = false
	currentColor = NewColor(0)
	currentTime  = NewSunriseSunset()
	neopixel     *ws2811.WS2811
)

func main() {
	if neopixel != nil {
		defer neopixel.Fini()
	}

	router := mux.NewRouter()

	router.HandleFunc("/color", getColor).Methods(http.MethodGet)
	router.HandleFunc("/color", postColor).Methods(http.MethodPost)

	router.HandleFunc("/time", getTime).Methods(http.MethodGet)

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for t := range ticker.C {
			if t.Day() != currentTime.Sunrise.Day() {
				currentTime = NewSunriseSunset()
			}
			if t.After(currentTime.Sunrise) && t.Before(currentTime.Sunset) {
				if ledsOn {
					go unsetColor()
					ledsOn = false
				}
			} else {
				if !ledsOn {
					go setColor(currentColor.Hex)
					ledsOn = true
				}
			}
		}
	}()

	log.Println("Listening on port 8040...")
	log.Fatal(http.ListenAndServe(":8040", router))
}

func getColor(w http.ResponseWriter, req *http.Request) {
	log.Println("GET /color")

	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(currentColor)
}

func postColor(w http.ResponseWriter, req *http.Request) {
	log.Println("POST /color")
	json.NewDecoder(req.Body).Decode(currentColor)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(currentColor)

	if ledsOn {
		go setColor(currentColor.Hex)
	}
}

func getTime(w http.ResponseWriter, req *http.Request) {
	log.Println("GET /time")

	w.Header().Add("Access-Control-Allow-Origin", "*")
	time := NewTime(currentTime.Sunrise, currentTime.Sunset)
	json.NewEncoder(w).Encode(time)
}
