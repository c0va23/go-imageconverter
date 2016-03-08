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
	cmd *exec.Cmd,
	imageData []byte,
	width, height uint,
) (
	[]byte,
	error,
) {
	cmd.Args = append(cmd.Args, convertArgs(width, height)...)
	cmd.Stdin = bytes.NewBuffer(imageData)

	convertedImageData, runErr := cmd.Output()
	if nil != runErr {
		log.Printf("Error output: %s", runErr)
	}
	return convertedImageData, runErr
}

func imageMagickCmdConvertImage(imageData []byte, width, height uint) ([]byte, error) {
	cmd := exec.Command("convert")
	return convertCmdConvertImage(cmd, imageData, width, height)
}

func graphicsmagickCmdConvertImage(imageData []byte, width, height uint) ([]byte, error) {
	cmd := exec.Command("gm", "convert")
	return convertCmdConvertImage(cmd, imageData, width, height)
}
