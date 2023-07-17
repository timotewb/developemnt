package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

var api_url string = "http://192.168.144.210:80/api"
var sdb_url string = "http://192.168.144.131:8000/sql"

func surreadDBCall(sdb_url string, api_url string, url string) {
	body := []byte(`insert into rss_news (
		select 
			string::join('-',URL,PubDate) as id,
			*
		from http::post(
			'` + api_url + `',
			{
				'name':'many-rss_xml',
				'args':[
					'` + url + `'
				]
			}
		));`)
	r, err := http.NewRequest("POST", sdb_url, bytes.NewBuffer(body))
	r.SetBasicAuth("etl", "etl")
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("NS", "many")
	r.Header.Add("DB", "db01")

	resp, err := myClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var j = []byte(`{"code":200,"details":"success","description":"All OK","information":"All OK"}`)
		fmt.Println(string(j))
	}
}

func main() {

	readFile, err := os.Open("/mnt/ns01/servers/factotum/01/api/apps/urls.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		if fileScanner.Text() != "" {
			surreadDBCall(sdb_url, api_url, fileScanner.Text())
			var d int = rand.Intn(500)
			time.Sleep(time.Duration(d) * time.Millisecond)
		}
	}

	readFile.Close()
}
