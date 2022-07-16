package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

func file_get_contents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func getCityCode(city string) {
	var c City
	content, err := file_get_contents("city.list.json")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		log.Panic(err)
	}
}

func getData(code string) {
	// TODO: get data from API
}

func main() {
	server()
}
