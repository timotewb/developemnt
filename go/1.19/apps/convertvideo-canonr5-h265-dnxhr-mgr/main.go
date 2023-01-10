package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var fileList []string
var myClient = &http.Client{Timeout: 10000 * time.Second}

var api_url string = "http://192.168.144.210:80/api"
var compute_url string = "http://192.168.144.210:90/api"
var dir string = "/mnt/ns00/Convert/CanonR5/h265/"

func main() {

	// setup call
	body := []byte(`{"Name":"listdir","Args":["` + dir + `"]}`)
	r, err := http.NewRequest("POST", api_url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")

	// call api
	resp, err := myClient.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// get response into []string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fileList = strings.Split(string(bodyBytes)[1:len(string(bodyBytes))-1], ",")
	}

	// call conversion api
	for i, f := range fileList {

		// setup call
		var outFile string = `"` + dir + strings.ReplaceAll(strings.TrimSpace(f), "'", "") + `"`
		fmt.Println("Running: ", i, f, `{"Name":"linux_amd64/convertvideo-canonr5-h265-dnxhr","Args":[`+outFile+`]}`)
		// setup call
		body := []byte(`{"Name":"linux_amd64/convertvideo-canonr5-h265-dnxhr","Args":[` + outFile + `]}`)
		r, err := http.NewRequest("POST", compute_url, bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}
		r.Header.Add("Content-Type", "application/json")

		// call api
		resp, err := myClient.Do(r)
		if err != nil {
			if os.IsTimeout(err) {
				// A timeout error occurred
				continue
			} else {
				panic(err)
			}
		}
		defer resp.Body.Close()

		// get response into []string
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(bodyBytes))
		}
	}
}
