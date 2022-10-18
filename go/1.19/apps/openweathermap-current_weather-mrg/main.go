package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}
var groupSize int = 7

type CityList []struct {
	Time   string `json:"time"`
	Status string `json:"status"`
	Result []struct {
		LocationID int `json:"location_id"`
	} `json:"result"`
}

func getJson(url string, target interface{}) {
	body := []byte("SELECT location_id FROM city_list limit 30;")

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	r.SetBasicAuth("root", "root")
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("NS", "openweathermap")
	r.Header.Add("DB", "db01")

	resp, err := myClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//bodyString := string(bodyBytes)
		//fmt.Println(bodyString)
		if err := json.Unmarshal(bodyBytes, &target); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
	}
}

func main() {

	// get list of city ids
	var cl CityList
	getJson("http://localhost:8000/sql", &cl)

	// break ids into groups and call
	var l int = 1
	var s string = ""
	for i, o := range cl[0].Result {
		//fmt.Println(i, o.LocationID)
		l += 1
		if s == "" {
			s = strconv.Itoa(o.LocationID)
		} else {
			s = s + "," + strconv.Itoa(o.LocationID)
		}
		if l == groupSize || i+1 == len(cl[0].Result) {
			//fmt.Println(s)
			// change this to call SurrealDB to call this endpoint
			cmd := exec.Command("apps/openweathermap-current_weather", "df128806bcff028c84dc038ccfcaaa44", s)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(string(stdout))
			l = 1
			s = ""
		}
	}
}
