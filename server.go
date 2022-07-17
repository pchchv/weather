package main

import (
	"fmt"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(getJSON("", "Weather Service. Version 0.2"))
	if err != nil {
		log.Panic(err)
	}
}

func cityWeather(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query().Get("city")
	city := getCityData(c)
	data := getData(city)
	weather := getWeather(data)
	resp := fmt.Sprintf("The temperature in %v is %v degrees Celsius.", c, weather)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(getJSON("", resp))
	if err != nil {
		log.Panic(err)
	}
}

func cityTime(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query().Get("city")
	city := getCityData(c)
	time := getTime(city)
	resp := fmt.Sprintf("Time in %v now: %v", c, time)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(getJSON("", resp))
	if err != nil {
		log.Panic(err)
	}
}

func cityStats(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query().Get("city")
	city := getCityData(c)
	weather := getWeather(getData(city))
	time := getTime(city)
	t := fmt.Sprintf("Time in %v now: %v", c, time)
	we := fmt.Sprintf("The temperature in %v is %v degrees Celsius.", c, weather)
	resp := getJSON("", t)
	resp = append(resp, getJSON("", we)...)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resp)
	if err != nil {
		log.Panic(err)
	}
}

func server() {
	log.Println("Server started!")
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/weather", cityWeather)
	http.HandleFunc("/time", cityTime)
	http.HandleFunc("/stats", cityStats)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
