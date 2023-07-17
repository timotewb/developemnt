package main

import (
	"log"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	err := ffmpeg.Input("in1.mp4").
		Output("out1.mp4", ffmpeg.KwArgs{"c:v": "libx265"}).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Fatal(err)
	}
}
