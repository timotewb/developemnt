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
	List []Response `json:list,omitempty`
}

type Response struct {
	Coord      Coordinates      `json:coord,omitempty`
	Weather    []WeatherDetails `json:weather,omitempty`
	Base       string           `json:base,omitempty`
	Main       MainDetails      `json:main,omitempty`
	Visibility int32            `json:visibility,omitempty`
	Wind       WindDetails      `json:wind,omitempty`
	Clouds     CloudDetails     `json:clouds,omitempty`
	Rain       RainDetails      `json:rain,omitempty`
	Snow       SnowDetails      `json:snow,omitempty`
	Dt         int32            `json:dt,omitempty`
	Sys        SysDetails       `json:sys,omitempty`
	Timezone   int32            `json:timezone,omitempty`
	Id         int32            `json:id,omitempty`
	Name       string           `json:name,omitempty`
	Cod        int16            `json:cod,omitempty`
}

type Coordinates struct {
	Lon float32 `json:lon,omitempty`
	Lat float32 `json:lat,omitempty`
}

type WeatherDetails struct {
	Id          int32  `json:id,omitempty`
	Main        string `json:main,omitempty`
	Description string `json:description,omitempty`
	Icon        string `json:icon,omitempty`
}

type MainDetails struct {
	Temp       float32 `json:temp,omitempty`
	Feels_like float32 `json:feels_like,omitempty`
	Pressure   int32   `json:pressure,omitempty`
	Humidity   int32   `json:humidity,omitempty`
	Temp_min   float32 `json:temp_min,omitempty`
	Temp_max   float32 `json:temp_max,omitempty`
	Sea_level  float32 `json:sea_level,omitempty`
	Grnd_level float32 `json:grnd_level,omitempty`
}

type WindDetails struct {
	Speed float32 `json:speed,omitempty`
	Deg   float32 `json:deg,omitempty`
}

type CloudDetails struct {
	All int16 `json:all,omitempty`
}

type RainDetails struct {
	Rain1h int16 `json:1h,omitempty`
	Rain3h int16 `json:3h,omitempty`
}

type SnowDetails struct {
	Snow1h int16 `json:1h,omitempty`
	Snow3h int16 `json:3h,omitempty`
}

type SysDetails struct {
	Type    int16  `json:type,omitempty`
	Id      int32  `json:id,omitempty`
	Message string `json:message,omitempty`
	Country string `json:country,omitempty`
	Sunrise int32  `json:sunrise,omitempty`
	Sunset  int32  `json:sunset,omitempty`
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
		//fmt.Println(data.Cnt)
		if data.Cnt > 0 {
			//b, err := json.MarshalIndent(data.List, "", "  ")
			b, err := json.Marshal(data.List)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(string(b)) // TODO find a way to lower case json key names
		}
	}

}
