package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// var url string = "http://pokeapi.co/api/v2/pokedex/kanto/"
var url string = "https://api.openweathermap.org/data/2.5/weather?id=2179537&appid=df128806bcff028c84dc038ccfcaaa44&units=metric&lang=en"

type Response struct {
	Coord      Coordinates      `json:coord`
	Weather    []WeatherDetails `json:weather`
	Base       string           `json:base`
	Main       MainDetails      `json:main`
	Visibility int32            `json:visibility`
	Wind       WindDetails      `json:wind`
	Clouds     CloudDetails     `json:wind`
	Rain       RainDetails      `json:wind`
	Snow       SnowDetails      `json:wind`
	Dt         int32            `json:dt`
	Sys        SysDetails       `json:sys`
	Timezone   int32            `json:timezone`
	Id         int32            `json:id`
	Name       string           `json:name`
	Cod        int16            `json:cod`
}

type Coordinates struct {
	Lon float32 `json:lon`
	Lat float32 `json:lat`
}

type WeatherDetails struct {
	Id          int32  `json:id`
	Main        string `json:main`
	Description string `json:description`
	Icon        string `json:icon`
}

type MainDetails struct {
	Temp       float32 `json:temp`
	Feels_like float32 `json:feels_like`
	Pressure   int32   `json:pressure`
	Humidity   int32   `json:humidity`
	Temp_min   float32 `json:temp_min`
	Temp_max   float32 `json:temp_max`
	Sea_level  float32 `json:sea_level`
	Grnd_level float32 `json:grnd_level`
}

type WindDetails struct {
	Speed float32 `json:speed`
	Deg   float32 `json:deg`
}

type CloudDetails struct {
	All int16 `json:all`
}

type RainDetails struct {
	Rain1h int16 `json:1h`
	Rain3h int16 `json:3h`
}

type SnowDetails struct {
	Snow1h int16 `json:1h`
	Snow3h int16 `json:3h`
}

type SysDetails struct {
	Type    int16  `json:type`
	Id      int32  `json:id`
	Message string `json:message`
	Country string `json:country`
	Sunrise int32  `json:sunrise`
	Sunset  int32  `json:sunset`
}

func main() {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var data Response
	json.Unmarshal(responseData, &data)
	if data.Cod == 200 {
		fmt.Println(" --- struct ---")
		fmt.Println(data)
	}

}
