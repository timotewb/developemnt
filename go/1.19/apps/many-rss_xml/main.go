package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var deBug bool = false

// var getURL string = "https://rss.nytimes.com/services/xml/rss/nyt/World.xml"
// var getURL string = "https://www.rnz.co.nz/rss/business.xml"

// input struct(s)
type DataFormat struct {
	XMLName xml.Name `xml:"rss"`
	Rss     Channel  `xml:"channel"`
}
type Channel struct {
	// XMLName xml.Name `xml:"channel"`
	Items []Item `xml:"item"`
}
type Item struct {
	// XMLName     xml.Name     `xml:"item"`
	Title       string       `xml:"title"`
	Link        string       `xml:"link"`
	Description string       `xml:"description"`
	Creator     string       `xml:"creator"`
	PubDate     string       `xml:"pubDate"`
	ImageURL    ImageURLAttr `xml:"content"`
	ImageCredit string       `xml:"credit"`
}
type ImageURLAttr struct {
	// XMLName xml.Name `xml:"content"`
	Url string `xml:"url,attr"`
}

// output struct
type DataFormatOut struct {
	Items []ItemsOut
}
type ItemsOut struct {
	Title       string
	URL         string
	Description string
	Creator     string
	PubDate     string
	ImageURL    string
	ImageCredit string
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

	// define var types
	var agr1 string = os.Args[1]

	if xmlBytes, err := getXML(agr1); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		// remove any difficult strings
		xmlBytes = []byte(strings.ReplaceAll(string(xmlBytes), "atom:link", "atomlink"))

		var result DataFormat
		xml.Unmarshal(xmlBytes, &result)
		if err := xml.Unmarshal(xmlBytes, &result); err != nil {
			log.Fatal(err)
		} else {
			if deBug {
				for i, s := range result.Rss.Items {
					fmt.Println(i)
					fmt.Printf("Title:       %v\n", s.Title)
					fmt.Printf("Link:        %v\n", s.Link)
					fmt.Printf("Description: %v\n", strings.TrimSpace(s.Description))
					fmt.Printf("Creator:     %v\n", s.Creator)
					fmt.Printf("PubDate:     %v\n", s.PubDate)
					fmt.Printf("ImageURL:    %v\n", s.ImageURL)
					fmt.Printf("ImageCredit: %v\n", s.ImageCredit)
					fmt.Println()
				}
			}

			// tidy up and output to file as an array of structs (list of json)
			var outData []ItemsOut
			for _, d := range result.Rss.Items {
				var o ItemsOut
				o.Creator = strings.TrimSpace(d.Creator)
				o.Description = strings.TrimSpace(d.Description)
				o.ImageCredit = strings.TrimSpace(d.ImageCredit)
				o.ImageURL = strings.TrimSpace(d.ImageURL.Url)
				o.URL = strings.TrimSpace(d.Link)
				o.PubDate = strings.TrimSpace(d.PubDate)
				o.Title = strings.TrimSpace(d.Title)
				outData = append(outData, o)
			}
			b, err := json.MarshalIndent(outData, "", " ")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Print(string(b))
			//_ = ioutil.WriteFile("test.json", b, 0777)
		}
	}
}
