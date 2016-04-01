// +build vips

package main

import (
	"log"

	bimg "gopkg.in/h2non/bimg.v0"
)

const converterName = "vips"

func converterInitialize() {}

func converterTerminate() {}

func converter(imageData []byte, width, height uint) (output []byte, err error) {
	image := bimg.NewImage(imageData)

	resizedImageData, resizeErr := image.Resize(int(width), int(height))
	if nil != resizeErr {
		log.Printf("Error resize: %s", resizeErr)

		return nil, resizeErr
	}

	return resizedImageData, nil
}
