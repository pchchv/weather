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
	"strings"

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

type Weather struct {
	City        string
	Temperature float64
	Humidity    float64
	Pressure    float64
	Clouds      string
	WindSpeed   float64
}

type Time struct {
	Date string
	Time string
}

type Statistic struct {
	Weather Weather
	Time    Time
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

func file_get_contents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
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

func getWeatherData(city City) map[string]interface{} {
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

func getWeather(city string) []byte {
	var temp float64
	var humidity float64
	var pressure float64
	var clouds string
	var windSpeed float64
	cityData := getCityData(city)
	weather := getWeatherData(cityData)
	for i, j := range weather {
		switch i {
		case "main":
			w := j.(map[string]interface{})
			temp = w["temp"].(float64)
			humidity = w["temp"].(float64)
			pressure = w["pressure"].(float64)
		case "weather":
			w := j.([]interface{})[0].(map[string]interface{})
			clouds = w["description"].(string)
		case "wind":
			w := j.(map[string]interface{})
			windSpeed = w["speed"].(float64)
		}
	}
	resWeather := Weather{city, temp, humidity, pressure, clouds, windSpeed}
	w, err := json.MarshalIndent(resWeather, "  ", "\t")
	if err != nil {
		log.Panic(err)
	}
	return w
}

func getTimeData(city City) string {
	var data map[string]interface{}
	url := fmt.Sprintf("http://api.geonames.org/timezoneJSON?lat=%v&lng=%v&username=%v", city.Coord.Latitude, city.Coord.Longitude, getEnvValue("APIUSER"))
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
	return data["time"].(string)
}

func getTime(city string) []byte {
	cityData := getCityData(city)
	dt := strings.Split(getTimeData(cityData), " ")
	date := dt[0]
	time := dt[1]
	t := Time{date, time}
	d, err := json.MarshalIndent(t, "  ", "\t")
	if err != nil {
		log.Panic(err)
	}
	return d
}

func prnt(w map[string]interface{}) {
	for i, j := range w {
		fmt.Println(i)
		fmt.Println(j)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func main() {
	server()
}
