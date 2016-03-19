// +build !cmd,gm

package main

import (
	"fmt"
	"log"
	"unsafe"
)

// #cgo pkg-config: GraphicsMagickWand
// #include <stdlib.h>
// #include <wand/magick_wand.h>
import "C"

const converterName = "graphicsmagick"

func converterInitialize() {
	C.InitializeMagick(nil)
}

func converterTerminate() {
	C.DestroyMagick()
}

func converter(imageDate []byte, width, height uint) (output []byte, err error) {
	magickWand := C.NewMagickWand()
	defer C.DestroyMagickWand(magickWand)

	inBlob := unsafe.Pointer(&imageDate[0])
	inLenght := len(imageDate)

	if readBlobResult := C.MagickReadImageBlob(
		magickWand,
		(*C.uchar)(inBlob),
		C.size_t(inLenght),
	); C.MagickPass != readBlobResult {
		// TODO: Use C.MagickGetException for full message
		return nil, fmt.Errorf("Error read blob")
	}

	if scaleImageResult := C.MagickScaleImage(
		magickWand,
		C.ulong(width),
		C.ulong(height),
	); C.MagickPass != scaleImageResult {
		// TODO: Use C.MagickGetException for full message
		return nil, fmt.Errorf("Error scale image")
	}

	var outLength C.size_t
	outBlob := unsafe.Pointer(C.MagickWriteImageBlob(magickWand, &outLength))
	defer C.free(outBlob)

	outData := C.GoBytes(outBlob, C.int(outLength))

	log.Printf("Out image size: %d", len(outData))

	return outData, nil
}
