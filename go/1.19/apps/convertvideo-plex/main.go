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

var inDir string = "/mnt/ns00/Convert/Plex/"
var outDir string = "/mnt/ns00/Complete/Plex/"
var plexDirs [2]string = [2]string{"Movies", "TV Shows"}
var myClient = &http.Client{Timeout: 10 * time.Second}
var api_url string = "http://192.168.144.210:80/api"
var fileList []string

func main() {

	// for each folder
	for _, dir := range plexDirs {
		fmt.Println(dir)
		// get files
		body := []byte(`{"Name":"listdir","Args":["` + filepath.Join(inDir, dir) + `"]}`)
		r, err := http.NewRequest("POST", api_url, bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}
		r.Header.Add("Content-Type", "application/json")
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
			fileList = strings.Split(string(bodyBytes)[1:len(string(bodyBytes))-1], ",")
		}

		// create out dir if not exist
		if _, err := os.Stat(filepath.Join(outDir, dir)); os.IsNotExist(err) {
			err := os.Mkdir(filepath.Join(outDir, dir), 0777)
			if err != nil {
				log.Fatal(err)
			}
		}

		// for each file
		for _, f := range fileList {

			if f != "" {

				inFileDir := filepath.Join(inDir, dir, strings.TrimSpace(strings.ReplaceAll(f, "'", "")))
				fileName := strings.ReplaceAll(strings.TrimSpace(f[:len(f)-len(filepath.Ext(f))]), "'", "")
				outFIleDIr := filepath.Join(outDir, dir, fileName)

				fmt.Println("-", inFileDir)

				// convert file if not exist
				if _, err := os.Stat(filepath.Join(outDir, fileName+`.mov`)); errors.Is(err, os.ErrNotExist) {

					//HandBrakeCLI -v 1 -i "/mnt/ns00/Convert/Family Guy s04e30.mp4" -o "/mnt/ns00/Drop/Family Guy s04e30.m4v" -Z "Apple 2160p60 4K HEVC Surround" --encoder nvenc_h265
					cmd := exec.Command("HandBrakeCLI", "-v", "1", "-i", inFileDir, "-o", outFIleDIr+".m4v", "-Z", "Apple 2160p60 4K HEVC Surround", "--encoder", "nvenc_h265")
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

	}

	// shutdown
	cmd2 := exec.Command("/bin/sh", "-c", "sudo shutdown now")
	var out2 bytes.Buffer
	var stderr2 bytes.Buffer
	cmd2.Stdout = &out2
	cmd2.Stderr = &stderr2
	err := cmd2.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr2.String())
		return
	}
}
