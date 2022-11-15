package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}
var groupSize int = 20

var api_url string = "http://192.168.144.210:80/api"
var sdb_url string = "http://192.168.144.131:8000/sql"

// var api_url string = "http://localhost:3000/api"
// var sdb_url string = "http://localhost:8000/sql"

type CityList []struct {
	Time   string `json:"time"`
	Status string `json:"status"`
	Result []struct {
		LocationID int `json:"location_id"`
	} `json:"result"`
}

func getJson(sdb_url string, target interface{}) {
	body := []byte("SELECT location_id FROM city_list;")

	r, err := http.NewRequest("POST", sdb_url, bytes.NewBuffer(body))
	r.SetBasicAuth("etl", "etl")
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
		if err := json.Unmarshal(bodyBytes, &target); err != nil {
			fmt.Println("getJson - Can not unmarshal JSON")
		}
	}
}

type Body struct {
	Code        int    `json:"code"`
	Details     string `json:"details"`
	Description string `json:"description"`
	Information string `json:"information"`
}

func surreadDBCall(sdb_url string, api_url string, ids string, target interface{}) {
	body := []byte(`insert into current_weather (
		select 
			string::join('-',id,dt) as id,
			id as location_id,
			*
		from http::post(
			'` + api_url + `',
			{
				'name':'openweathermap-current_weather',
				'args':[
					'df128806bcff028c84dc038ccfcaaa44',
					'` + ids + `'
				]
			}
		));`)
	r, err := http.NewRequest("POST", sdb_url, bytes.NewBuffer(body))
	r.SetBasicAuth("etl", "etl")
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

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(bodyBytes, &target); err != nil {
			fmt.Println("surreadDBCall - Can not unmarshal JSON", err)
			fmt.Println(string(bodyBytes))
		}
	} else {
		var j = []byte(`{"code":200,"details":"success","description":"All OK","information":"All OK"}`)
		if err := json.Unmarshal(j, &target); err != nil {
			fmt.Println("surreadDBCall - Can not unmarshal JSON", err)
			fmt.Println(string(j))
		}
	}

}

func main() {

	// get list of city ids
	var cl CityList
	getJson(sdb_url, &cl)

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
			// cmd := exec.Command("apps/openweathermap-current_weather", "df128806bcff028c84dc038ccfcaaa44", s)
			// stdout, err := cmd.Output()
			// if err != nil {
			// 	fmt.Println(err.Error())
			// 	return
			// }
			// fmt.Println(string(stdout))

			// call SurrealDB to get data
			var b Body
			surreadDBCall(sdb_url, api_url, s, &b)
			if b.Code != 200 {
				b, err := json.Marshal(b)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(b))
				return
			}
			l = 1
			s = ""
			var d int = rand.Intn(500)
			time.Sleep(time.Duration(d) * time.Millisecond)
		}
	}
	fmt.Println("Manager job complete.")
}
