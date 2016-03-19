// +build !cmd,gm

package main

import (
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
	readBlobResult := C.MagickReadImageBlob(magickWand, (*C.uchar)(inBlob), C.size_t(inLenght))
	log.Printf("MagickReadImageBlob: %v", readBlobResult)

	scaleImageResult := C.MagickScaleImage(magickWand, C.ulong(width), C.ulong(height))
	log.Printf("MagickScaleImage: %v", scaleImageResult)

	var outLength C.size_t
	outBlob := unsafe.Pointer(C.MagickWriteImageBlob(magickWand, &outLength))
	defer C.free(outBlob)

	outData := C.GoBytes(outBlob, C.int(outLength))

	log.Printf("Out image size: %d", len(outData))

	return outData, nil
}
