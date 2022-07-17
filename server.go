package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := json.MarshalIndent("Weather Service. Version 1.0", "\t", "\t")
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(res)
	if err != nil {
		log.Panic(err)
	}
}

func cityWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	weather := getWeather(city)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(weather)
	if err != nil {
		log.Panic(err)
	}
}

func cityTime(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	time := getTime(city)
	res, err := json.Marshal(city + ": ")
	if err != nil {
		log.Panic(err)
	}
	res = append(res, time...)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		log.Panic(err)
	}
}

func cityStats(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	weather := getWeather(city)
	time := getTime(city)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(append(time, weather...))
	if err != nil {
		log.Panic(err)
	}
}

func server() {
	mux := http.NewServeMux()
	log.Println("Server started!")
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/weather", cityWeather)
	mux.HandleFunc("/time", cityTime)
	mux.HandleFunc("/stats", cityStats)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
