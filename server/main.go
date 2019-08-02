package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/color", getColor).Methods(http.MethodGet)
	router.HandleFunc("/color", postColor).Methods(http.MethodPost)

	router.HandleFunc("/time", getTime).Methods(http.MethodGet)
	router.HandleFunc("/time", postTime).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8040", router))
}

func getColor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("GET /color", params)
}

func postColor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("POST /color", params)
}

func getTime(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("GET /time", params)
}

func postTime(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Println("POST /time", params)
}
