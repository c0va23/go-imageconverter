// +build cmd

package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func converterInitialize() {}

func converterTerminate() {}

func convertArgs(width, heigth uint) []string {
	return []string{
		"-",
		"-scale", fmt.Sprintf("%dx%d", width, heigth),
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
