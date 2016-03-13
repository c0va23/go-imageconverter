// +build cmd,im

package main

import (
	"os/exec"
)

const converterName = "cmd,im"

func converter(imageData []byte, width, height uint) ([]byte, error) {
	cmd := exec.Command("convert")
	return convertCmdConvertImage(cmd, imageData, width, height)
}
