package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// SunriseSunset ...
type SunriseSunset struct {
	Results `json:"results"`
}

// Results ...
type Results struct {
	Sunrise time.Time `json:"sunrise"`
	Sunset  time.Time `json:"sunset"`
}

// NewSunriseSunset ...
func NewSunriseSunset() *SunriseSunset {
	sunriseSunset := &SunriseSunset{}

	resp, err := http.Get(fmt.Sprintf("https://api.sunrise-sunset.org/json?lat=%f&lng=%f&formatted=0", lat, lng))
	if err != nil {
		// TODO: Handle error
		panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(sunriseSunset)
	if err != nil {
		// TODO: Handle error
		panic(err)
	}

	location, err := time.LoadLocation(loc)
	if err != nil {
		// TODO: Handle error
		panic(err)
	}

	sunriseSunset.Sunrise = sunriseSunset.Sunrise.In(location).Add(1 * time.Hour)
	sunriseSunset.Sunset = sunriseSunset.Sunset.In(location).Add(-1 * time.Hour)

	log.Println(sunriseSunset)
	return sunriseSunset
}
