package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"gopkg.in/gographics/imagick.v2/imagick"
)

const imagesDir = "data/"

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
	return ioutil.ReadFile(imagesDir + imagePath)
}

func magickWandConvertImage(imageData []byte) ([]byte, error) {
	magickWand := imagick.NewMagickWand()
	defer magickWand.Destroy()

	if readErr := magickWand.ReadImageBlob(imageData); nil != readErr {
		log.Printf("Error read: %s", readErr)
		return nil, readErr
	}

	if resizeErr := magickWand.ResizeImage(800, 600, imagick.FILTER_LANCZOS, 1.0); nil != resizeErr {
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
				"-resize",
				"800x600",
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

const (
	converterMagickwand     = "magickwand"
	converterImageMagick    = "imagemagick"
	converterGraphicsMagick = "graphicsmagick"
)

func init() {
	flag.StringVar(
		&converterName,
		"converter",
		"magickwand",
		fmt.Sprintf(
			"Converter: %v",
			[]string{
				converterMagickwand,
				converterImageMagick,
				converterGraphicsMagick,
			},
		),
	)
	flag.Parse()

	flag.StringVar(&listen, "listen", "[::]:5050", "Listen address ip:port")

	log.Printf("Args: %v", os.Args)

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
