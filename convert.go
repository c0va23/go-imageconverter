package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func convertArgs(width, heigth uint) []string {
	return []string{
		"-",
		"-resize", fmt.Sprintf("%dx%d", width, heigth),
		"-",
	}
}

func convertCmdConvertImage(
	command string,
	imageData []byte,
	width, height uint,
	prefixArgs ...string,
) (
	[]byte,
	error,
) {
	cmd := exec.Command(
		command,
		append(
			prefixArgs,
			convertArgs(width, height)...,
		)...,
	)

	cmd.Stdin = bytes.NewBuffer(imageData)

	convertedImageData, runErr := cmd.Output()
	if nil != runErr {
		log.Printf("Error output: %s", runErr)
	}
	return convertedImageData, runErr
}

func imageMagickCmdConvertImage(imageData []byte, width, height uint) ([]byte, error) {
	return convertCmdConvertImage("convert", imageData, width, height)
}

func graphicsmagickCmdConvertImage(imageData []byte, width, height uint) ([]byte, error) {
	return convertCmdConvertImage("gm", imageData, width, height, "convert")
}
