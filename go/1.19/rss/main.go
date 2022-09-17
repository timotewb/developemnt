package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DataFormat struct {
	Rss struct {
		Channel struct {
			Title string `xml:title json:title`
		} `xml:channel json:channel`
	} `xml:rss json:rss`
}

// tweaked from: https://stackoverflow.com/a/42718113/1170664
func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func main() {
	if xmlBytes, err := getXML("https://rss.nytimes.com/services/xml/rss/nyt/World.xml"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		fmt.Print(string(xmlBytes))
		var result DataFormat
		xml.Unmarshal(xmlBytes, &result)
		if err := xml.Unmarshal(xmlBytes, &result); err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Rss.Channel.Title)
	}
}
