package main

import (
	"bytes"
	"log"

	"github.com/rainycape/magick"
)

func init() {
	log.Printf("Magick use %s backend", magick.Backend())
}

func magickConverter(imageDate []byte, width, height uint) (output []byte, err error) {
	image, decodeErr := magick.DecodeData(imageDate)

	if nil != decodeErr {
		log.Printf("Error decode: %s", decodeErr)
		return nil, decodeErr
	}
	log.Println("Image decoded")

	resizedImage, resizeErr := image.Resize(int(width), int(height), magick.FMitchell)
	if nil != resizeErr {
		log.Printf("Error resize %s", resizeErr)
		return nil, resizeErr
	}
	log.Println("Image resized")

	outBuffer := new(bytes.Buffer)

	outInfo := magick.NewInfo()
	outInfo.SetFormat("jpeg")
	// outInfo.SetQuality(75)

	encodeErr := resizedImage.Encode(outBuffer, outInfo)
	if nil != encodeErr {
		log.Printf("Error encode: %s", encodeErr)
		return nil, encodeErr
	}

	return outBuffer.Bytes(), nil
}
