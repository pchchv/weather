package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type City struct {
	Id      float64     `json:"id"`
	Name    string      `json:"name"`
	State   string      `json:"state"`
	Country string      `json:"country"`
	Coord   Coordinates `json:"coord"`
}

type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
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

func getCityData(city string) City {
	var res City
	var c []City
	content, err := file_get_contents("city.list.json")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		log.Panic(err)
	}
	for _, v := range c {
		if v.Name == city {
			res = v
			break
		}
	}
	if res.Id == 0 {
		log.Panic(errors.New("City not found"))
	}
	return res
}

func getData(city City) map[string]interface{} {
	var data map[string]interface{}
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?units=metric&lat=%v&lon=%v&appid=%v", city.Coord.Latitude, city.Coord.Longitude, getEnvValue("APIKEY"))
	res, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Panic(errors.New("status not OK"))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panic(err)
		}
	}(res.Body)
	if err != nil {
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Panic(err)
	}
	return data
}

func getWeather(data map[string]interface{}) float64 {
	data = data["main"].(map[string]interface{})
	return data["temp"].(float64)
}

func main() {
	server()
}
