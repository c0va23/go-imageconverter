package main

import (
	"log"

	"gopkg.in/gographics/imagick.v2/imagick"
)

var magickWandConvertImage Converter = func(
	imageData []byte,
	width, heigth uint,
) (
	[]byte,
	error,
) {
	magickWand := imagick.NewMagickWand()
	defer magickWand.Destroy()

	if readErr := magickWand.ReadImageBlob(imageData); nil != readErr {
		log.Printf("Error read: %s", readErr)
		return nil, readErr
	}

	if scaleErr := magickWand.ScaleImage(
		width,
		heigth,
	); nil != scaleErr {
		log.Printf("Error scale: %s", scaleErr)
		return nil, scaleErr
	}

	convertedImage := magickWand.GetImageBlob()

	return convertedImage, nil
}
