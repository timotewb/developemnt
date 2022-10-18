package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ResponseWrapper struct {
	Cnt  int        `json:cnt`
	List []Response `json:list`
}

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

	// define var types
	var agr1 string = os.Args[1]
	var agr2 string = os.Args[2]
	//fmt.Println("https://api.openweathermap.org/data/2.5/group?id=" + agr2 + "&appid=" + agr1)

	// make call to api
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/group?id=" + agr2 + "&appid=" + agr1)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Error: Non 200 status code returned when attempting to retrieve file. Status Code was %v.\n", resp.StatusCode)
		os.Exit(1)
	}

	// convert respose to string then return
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		// return to calling api job
		//fmt.Print(string(bodyBytes))
		var data ResponseWrapper
		json.Unmarshal(bodyBytes, &data)
		fmt.Println(data.Cnt)
		if data.Cnt > 0 {
			b, err := json.MarshalIndent(data.List, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(b))
		}
	}

}
