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
	brightness = 128
)

var (
	currentColor = NewColor(0)
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
	router.HandleFunc("/time", postTime).Methods(http.MethodPost)

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

	go setColor(currentColor.Hex)
}

func getTime(w http.ResponseWriter, req *http.Request) {
	log.Println("GET /time")

	w.Header().Add("Access-Control-Allow-Origin", "*")
	time := NewTime(time.Now(), time.Now().Add(time.Hour))
	json.NewEncoder(w).Encode(time)
}

func postTime(w http.ResponseWriter, req *http.Request) {
	log.Println("POST /time")
}
