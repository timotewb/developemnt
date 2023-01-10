package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	// define var types
	var agr1 string = os.Args[1]
	//var agr1 string = "/mnt/ns00/Convert/CanonR5/h265/0B7A2766.MP4"

	// create required vars
	fileNameX := filepath.Base(agr1)
	outDir := strings.ReplaceAll(filepath.Dir(agr1), "Convert", "Complete")
	fileName := fileNameX[:len(fileNameX)-len(filepath.Ext(fileNameX))]
	//ffmpeg := `/usr/bin/ffmpeg -i ` + agr1 + ` -c:v dnxhd -vf "scale=3840:2160,fps=30000/1001,format=yuv422p10le" -profile:v dnxhr_hqx -b:v 873M -c:a pcm_s16le ` + filepath.Join(outDir, fileName+`.mov`)

	// create out dir if not exist
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err := os.Mkdir(outDir, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	// convert file if not exist
	if _, err := os.Stat(filepath.Join(outDir, fileName+`.mov`)); errors.Is(err, os.ErrNotExist) {

		cmd := exec.Command("/usr/bin/ffmpeg", "-i", agr1, "-c:v", "dnxhd", "-vf", "scale=3840:2160,fps=30000/1001,format=yuv422p10le", "-profile:v", "dnxhr_hqx", "-b:v", "873M", "-c:a", "pcm_s16le", filepath.Join(outDir, fileName+`.mov`))
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}
		fmt.Println("Result: " + out.String())
	} else {
		println("File already exists.")
	}

}
