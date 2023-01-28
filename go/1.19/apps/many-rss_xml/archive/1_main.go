package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	xj "github.com/basgys/goxml2json"
)

var getURL string = "https://rss.nytimes.com/services/xml/rss/nyt/World.xml"

//var getURL string = "https://www.rnz.co.nz/rss/business.xml"

// input struct(s)
type DataFormat struct {
	XMLName xml.Name `xml:"rss"`
	Rss     Channel  `xml:"channel"`
}
type Channel struct {
	// XMLName xml.Name `xml:"channel"`
	Items string `xml:"item"`
}

func main() {

	// make call to api
	resp, err := http.Get(getURL)
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

		// xml is an io.Reader
		xml := strings.NewReader(string(bodyBytes))
		json, err := xj.Convert(xml)
		if err != nil {
			panic("That's embarrassing...")
		}

		fmt.Println(json)
	}
}
