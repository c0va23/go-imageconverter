package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/gographics/imagick.v2/imagick"
)

const imagesDir = "data/"

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	imagePath := request.RequestURI[1:]
	log.Printf("Request image with path: %s\n", imagePath)

	if imageData, fetchErr := fetchOriginImage(imagePath); nil != fetchErr {
		log.Printf("Error fetch image date: %s", fetchErr)
		responseWriter.WriteHeader(http.StatusNotFound)
	} else if convertedImage, convertErr := convertImage(imageData); nil != convertErr {
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

func convertImage(imageData []byte) ([]byte, error) {
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

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	serverErr := http.ListenAndServe("[::]:5050", http.HandlerFunc(handler))
	if nil != serverErr {
		panic(serverErr)
	}
}
