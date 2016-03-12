package main

import (
	"flag"
	"fmt"
	"log"
)

// Command line options
var converterName string
var converter Converter
var listen string
var imagesDir string
var outWidth uint
var outHeigth uint

const (
	converterMagickwand     = "magickwand"
	converterImageMagick    = "imagemagick"
	converterGraphicsMagick = "graphicsmagick"
	converterMagick         = "magick"
)

var converters = []string{
	converterMagickwand,
	converterImageMagick,
	converterGraphicsMagick,
	converterMagick,
}

func init() {
	flag.StringVar(&converterName, "converter", converterMagickwand,
		fmt.Sprintf("Converter: %v", converters))

	flag.StringVar(&listen, "listen", ":5050", "Listen address ip:port")
	flag.StringVar(&imagesDir, "images-dir", "data", "Images root directory")
	flag.UintVar(&outWidth, "out-width", 800, "Out width")
	flag.UintVar(&outHeigth, "out-heigth", 600, "Out heigth")

	flag.Parse()

	findConverter()

	log.Printf("Used converter: %s", converterName)
	log.Printf("Images root dir: %s", imagesDir)
	log.Printf("Out width x heigth: %d x %d", outWidth, outHeigth)
}

func findConverter() {
	switch converterName {
	case converterMagickwand:
		converter = magickWandConvertImage
	case converterImageMagick:
		converter = imageMagickCmdConvertImage
	case converterGraphicsMagick:
		converter = graphicsmagickCmdConvertImage
	case converterMagick:
		converter = magickConverter
	default:
		log.Fatalf("Invalid converter: %s", converterName)
	}
}
