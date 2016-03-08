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

	if resizeErr := magickWand.ResizeImage(
		width,
		heigth,
		imagick.FILTER_LANCZOS,
		1.0,
	); nil != resizeErr {
		log.Printf("Error resize: %s", resizeErr)
		return nil, resizeErr
	}

	convertedImage := magickWand.GetImageBlob()

	return convertedImage, nil
}
