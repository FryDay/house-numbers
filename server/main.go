package main

import (
	"encoding/json"
	"flag"
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
	config       = new(Config)
	currentColor = NewColor(0)
	currentTime  = NewSunriseSunset()
	forceOn      = NewForceOn()
	ledsOnChan   = make(chan bool)
	neopixel     *ws2811.WS2811
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config/dev.toml", "config path")
	flag.Parse()
	config = NewConfig(configPath)

	if neopixel != nil {
		defer neopixel.Fini()
		unsetColor()
	}

	router := mux.NewRouter()

	router.HandleFunc("/color", getColor).Methods(http.MethodGet)
	router.HandleFunc("/color", postColor).Methods(http.MethodPost)

	router.HandleFunc("/force", getForceOn).Methods(http.MethodGet)
	router.HandleFunc("/force", postForceOn).Methods(http.MethodPost)

	router.HandleFunc("/time", getTime).Methods(http.MethodGet)

	go func() {
		ticker := time.NewTicker(1 * time.Minute)

		for t := range ticker.C {
			if t.Day() != currentTime.Sunrise.Day() {
				currentTime = NewSunriseSunset()
			}

			ledsOnChan <- (t.Before(currentTime.Sunrise) && t.After(currentTime.Sunset))
		}
	}()

	go func() {
		on := false
		lastHex := ""

		for {
			turnOn := <-ledsOnChan
			if turnOn {
				if !on ||
					(on && lastHex != currentColor.Hex) ||
					(!on && forceOn.On) {
					setColor(currentColor.Hex)
					lastHex = currentColor.Hex
					on = true
				}
			} else {
				if on && !forceOn.On {
					unsetColor()
					lastHex = ""
					on = false
				}
			}
		}
	}()

	log.Println("Listening on port 8040...")
	log.Fatal(http.ListenAndServe(":8040", router))
}

func getColor(w http.ResponseWriter, req *http.Request) {
	log.Println("GET /color")

	w.Header().Add("Access-Control-Allow-Origin", config.URL)
	json.NewEncoder(w).Encode(currentColor)
}

func getForceOn(w http.ResponseWriter, req *http.Request) {
	log.Println("GET /force")

	w.Header().Add("Access-Control-Allow-Origin", config.URL)
	json.NewEncoder(w).Encode(forceOn)
}

func postColor(w http.ResponseWriter, req *http.Request) {
	log.Println("POST /color")
	json.NewDecoder(req.Body).Decode(currentColor)

	w.Header().Add("Access-Control-Allow-Origin", config.URL)
	json.NewEncoder(w).Encode(currentColor)
}

func postForceOn(w http.ResponseWriter, req *http.Request) {
	log.Println("POST /force")

	w.Header().Add("Access-Control-Allow-Origin", config.URL)
	json.NewEncoder(w).Encode(forceOn)

	ledsOnChan <- forceOn.On
}

func getTime(w http.ResponseWriter, req *http.Request) {
	log.Println("GET /time")

	w.Header().Add("Access-Control-Allow-Origin", config.URL)
	time := NewTime(currentTime.Sunrise, currentTime.Sunset)
	json.NewEncoder(w).Encode(time)
}
