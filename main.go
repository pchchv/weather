package main

import (
	"encoding/json"
	"log"
)

type City struct {
	Id      uint64      `json:"id"`
	Name    string      `json:"name"`
	State   string      `json:"state"`
	Country string      `json:"country"`
	Coord   Coordinates `json:"coord"`
}

type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

func getJSON(pre string, str string) []byte {
	var res []byte
	s, err := json.MarshalIndent(str, "\t", "\t")
	if err != nil {
		log.Panic(err)
	}
	if pre != "" {
		pr, err := json.MarshalIndent(pre, "\t", "\t")
		if err != nil {
			log.Panic(err)
		}
		s = append(pr, s...)
	}
	res = append(res, s...)
	return res
}

func getCityCode(city string) string {
	var code string
	//TODO: Get data from JSON
	return code
}

func getData(code string) {
	// TODO: get data from API
}

func main() {
	server()
}
