package main

import (
	"log"
	"net/http"
)

func server() {
	log.Println("Server started!")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
