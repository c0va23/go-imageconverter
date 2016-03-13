// +build !cmd,gm

package main

import (
	"log"
	"unsafe"
)

/// # include <stdio.h>
/// # include <string.h>
/// # include <stdlib.h>

// #cgo pkg-config: GraphicsMagick
// #include <magick/api.h>
import "C"

const converterName = "graphicsmagick"

func converterInitialize() {
	C.InitializeMagick(nil)
}

func converterTerminate() {
	C.DestroyMagick()
}

func converter(imageDate []byte, width, height uint) (output []byte, err error) {
	var exception C.ExceptionInfo
	C.GetExceptionInfo(&exception)
	defer C.DestroyExceptionInfo(&exception)

	inImageInfo := C.CloneImageInfo(nil)
	defer C.DestroyImageInfo(inImageInfo)

	inBlob := unsafe.Pointer(&imageDate[0])
	inLenght := C.size_t(len(imageDate))
	inImage := C.BlobToImage(inImageInfo, inBlob, inLenght, &exception)
	C.CatchException(&exception)
	defer C.DestroyBlob(inImage)
	defer C.DestroyImage(inImage)

	outImage := C.ScaleImage(
		inImage,
		C.ulong(width),
		C.ulong(height),
		&exception,
	)
	C.CatchException(&exception)
	defer C.DestroyImage(outImage)

	outImageInfo := C.CloneImageInfo(inImageInfo)
	defer C.DestroyImageInfo(outImageInfo)

	var outLength C.size_t
	var outBlob unsafe.Pointer
	outBlob = C.ImageToBlob(outImageInfo, outImage, &outLength, &exception)
	C.CatchException(&exception)
	defer C.DestroyBlob(outImage)

	outData := C.GoBytes(outBlob, C.int(outLength))

	log.Printf("Out image size: %d", len(outData))

	return outData, nil
}
