package main

import (
	"flag"
	"log"
)

// Command line options
var listen string
var imagesDir string
var outWidth uint
var outHeigth uint

func init() {
	flag.StringVar(&listen, "listen", ":5050", "Listen address ip:port")
	flag.StringVar(&imagesDir, "images-dir", "data", "Images root directory")
	flag.UintVar(&outWidth, "out-width", 800, "Out width")
	flag.UintVar(&outHeigth, "out-heigth", 600, "Out heigth")

	flag.Parse()

	log.Printf("Images root dir: %s", imagesDir)
	log.Printf("Out width x heigth: %d x %d", outWidth, outHeigth)
}
