package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var fileList []string
var myClient = &http.Client{Timeout: 10 * time.Second}

var api_url string = "http://192.168.144.210:80/api"
var dir string = "/mnt/ns00/Convert/CanonR5/h265/"
var outDir string = "/mnt/ns00/Complete/CanonR5/h265/"

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

	// create out dir if not exist
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err := os.Mkdir(outDir, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	// convert files
	for _, f := range fileList {

		// create required vars
		fileNameX := f
		fileName := fileNameX[:len(fileNameX)-len(filepath.Ext(fileNameX))]
		//ffmpeg := `/usr/bin/ffmpeg -i ` + agr1 + ` -c:v dnxhd -vf "scale=3840:2160,fps=30000/1001,format=yuv422p10le" -profile:v dnxhr_hqx -b:v 873M -c:a pcm_s16le ` + filepath.Join(outDir, fileName+`.mov`)

		// convert file if not exist
		if _, err := os.Stat(filepath.Join(outDir, fileName+`.mov`)); errors.Is(err, os.ErrNotExist) {

			cmd := exec.Command("/usr/bin/ffmpeg", "-i", filepath.Join(outDir, fileNameX), "-c:v", "dnxhd", "-vf", "scale=3840:2160,fps=30000/1001,format=yuv422p10le", "-profile:v", "dnxhr_hqx", "-b:v", "873M", "-c:a", "pcm_s16le", filepath.Join(outDir, fileName+".mov"))
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
				return
			}
			fmt.Println("Result: Conversion complete. " + out.String())
		} else {
			println("Result: File already exists.")
		}

	}
}
