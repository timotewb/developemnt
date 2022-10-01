package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://bulk.openweathermap.org/sample/city.list.json.gz")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Error: Non 200 status code returned when attempting to retrieve file. Status Code was %v.\n", resp.StatusCode)
		os.Exit(1)
	}

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	var body []byte
	body, err = io.ReadAll(reader)
	reader.Close()

	fmt.Print(string(body))

}
