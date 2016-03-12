package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/gographics/imagick.v2/imagick"
)

// Converter is type function for convert image from imageDate to width x heigth
type Converter func(imageDate []byte, width, height uint) (output []byte, err error)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	imagePath := request.RequestURI[1:]
	log.Printf("Request image with path: %s\n", imagePath)

	if imageData, fetchErr := fetchOriginImage(imagePath); nil != fetchErr {
		log.Printf("Error fetch image date: %s", fetchErr)
		responseWriter.WriteHeader(http.StatusNotFound)
	} else if convertedImage, convertErr := converter(
		imageData,
		outWidth,
		outHeigth,
	); nil != convertErr {
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
