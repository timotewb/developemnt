package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DataFormat struct {
	XMLName xml.Name `xml:"rss"`
	Rss     Channel  `xml:"channel"`
}
type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []Item   `xml:"item"`
}
type Item struct {
	XMLName     xml.Name     `xml:"item"`
	Title       string       `xml:"title"`
	Link        string       `xml:"link"`
	Description string       `xml:"description"`
	Creator     string       `xml:"creator"`
	PubDate     string       `xml:"pubDate"`
	ImageURL    ImageURLAttr `xml:"content"`
	ImageCredit string       `xml:"credit"`
}
type ImageURLAttr struct {
	XMLName xml.Name `xml:"content"`
	Url     string   `xml:"url,attr"`
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
		//fmt.Print(string(xmlBytes))
		var result DataFormat
		xml.Unmarshal(xmlBytes, &result)
		if err := xml.Unmarshal(xmlBytes, &result); err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Rss.Items)
	}
}
