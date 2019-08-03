package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/color", getColor).Methods(http.MethodGet)
	router.HandleFunc("/color", postColor).Methods(http.MethodPost)

	router.HandleFunc("/time", getTime).Methods(http.MethodGet)
	router.HandleFunc("/time", postTime).Methods(http.MethodPost)

	log.Println("Listening on port 8040...")
	log.Fatal(http.ListenAndServe(":8040", router))
}

func getColor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("GET /color", params)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	// color := NewFromRGB(255, 255, 255)
	color := NewFromHex("#00ff00")
	json.NewEncoder(w).Encode(color)
}

func postColor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("POST /color", params)
}

func getTime(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("GET /time", params)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	time := NewTime(time.Now(), time.Now().Add(time.Hour))
	json.NewEncoder(w).Encode(time)
}

func postTime(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("POST /time", params)
}
