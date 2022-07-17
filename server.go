package main

import (
	"fmt"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(getJSON("", "Weather Service. Version 0.0.1"))
	if err != nil {
		log.Panic(err)
	}
}

func weather(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query().Get("city")
	city := getCityData(c)
	data := getData(city)
	weather := getWeather(data)
	resp := fmt.Sprintf("The temperature in %v is %v degrees Celsius.", c, weather)
	_, err := w.Write(getJSON("", resp))
	if err != nil {
		log.Panic(err)
	}
}

func server() {
	log.Println("Server started!")
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/weather", weather)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
