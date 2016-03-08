package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"gopkg.in/gographics/imagick.v2/imagick"
)

var converter func(imageDate []byte) (output []byte, err error)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	imagePath := request.RequestURI[1:]
	log.Printf("Request image with path: %s\n", imagePath)

	if imageData, fetchErr := fetchOriginImage(imagePath); nil != fetchErr {
		log.Printf("Error fetch image date: %s", fetchErr)
		responseWriter.WriteHeader(http.StatusNotFound)
	} else if convertedImage, convertErr := converter(imageData); nil != convertErr {
		log.Printf("Error convert image %s: %s", imagePath, convertErr)
		responseWriter.WriteHeader(http.StatusInternalServerError)
	} else if written, writeErr := responseWriter.Write(convertedImage); nil == writeErr {
		log.Printf("Image %s write %d bytes", imagePath, written)
	} else {
		log.Printf("Error write image %s: %s", imagePath, writeErr)
	}
}

func fetchOriginImage(imagePath string) ([]byte, error) {
	return ioutil.ReadFile(imagesDir + "/" + imagePath)
}

func magickWandConvertImage(imageData []byte) ([]byte, error) {
	magickWand := imagick.NewMagickWand()
	defer magickWand.Destroy()

	if readErr := magickWand.ReadImageBlob(imageData); nil != readErr {
		log.Printf("Error read: %s", readErr)
		return nil, readErr
	}

	if resizeErr := magickWand.ResizeImage(
		outWidth,
		outHeigth,
		imagick.FILTER_LANCZOS,
		1.0,
	); nil != resizeErr {
		log.Printf("Error resize: %s", resizeErr)
		return nil, resizeErr
	}

	convertedImage := magickWand.GetImageBlob()

	return convertedImage, nil
}

func convertCmdConvertImage(
	command string,
	imageData []byte,
	prefixArgs ...string,
) (
	[]byte,
	error,
) {
	cmd := exec.Command(
		command,
		append(
			prefixArgs,
			[]string{
				"-",
				"-resize", fmt.Sprintf("%dx%d", outWidth, outHeigth),
				"-",
			}...,
		)...,
	)

	cmd.Stdin = bytes.NewBuffer(imageData)

	convertedImageData, runErr := cmd.Output()
	if nil != runErr {
		log.Printf("Error output: %s", runErr)
	}
	return convertedImageData, runErr
}

func imageMagickCmdConvertImage(imageData []byte) ([]byte, error) {
	return convertCmdConvertImage("convert", imageData)
}

func graphicsmagickCmdConvertImage(imageData []byte) ([]byte, error) {
	return convertCmdConvertImage("gm", imageData, "convert")
}

// Command line options
var converterName string
var listen string
var imagesDir string
var outWidth uint
var outHeigth uint

const (
	converterMagickwand     = "magickwand"
	converterImageMagick    = "imagemagick"
	converterGraphicsMagick = "graphicsmagick"
)

var converters = []string{
	converterMagickwand,
	converterImageMagick,
	converterGraphicsMagick,
}

func init() {
	flag.StringVar(&converterName, "converter", converterMagickwand,
		fmt.Sprintf("Converter: %v", converters))

	flag.StringVar(&listen, "listen", ":5050", "Listen address ip:port")
	flag.StringVar(&imagesDir, "images-dir", "data", "Images root directory")
	flag.UintVar(&outWidth, "out-width", 800, "Out width")
	flag.UintVar(&outHeigth, "out-heigth", 600, "Out heigth")

	flag.Parse()

	switch converterName {
	case converterMagickwand:
		converter = magickWandConvertImage
	case converterImageMagick:
		converter = imageMagickCmdConvertImage
	case converterGraphicsMagick:
		converter = graphicsmagickCmdConvertImage
	default:
		log.Fatalf("Invalid converter: %s", converterName)
	}
	log.Printf("Used converter: %s", converterName)
	log.Printf("Images root dir: %s", imagesDir)
	log.Printf("Out width x heigth: %d x %d", outWidth, outHeigth)
}

func main() {
	if converterName == converterMagickwand {
		imagick.Initialize()
		defer imagick.Terminate()
	}

	log.Printf("Start listen on %s", listen)
	serverErr := http.ListenAndServe(listen, http.HandlerFunc(handler))
	if nil != serverErr {
		panic(serverErr)
	}
}
