package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://bulk.openweathermap.org/sample/city.list.json.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var body []byte
	body, err = io.ReadAll(reader)
	reader.Close()

	fmt.Printf(string(body))

}
