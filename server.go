package main

import (
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

func getTime(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
}

func server() {
	log.Println("Server started!")
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/time", getTime)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
