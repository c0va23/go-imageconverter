// +build cmd,gm

package main

import (
	"os/exec"
)

const converterName = "cmd,gm"

func converter(imageData []byte, width, height uint) ([]byte, error) {
	cmd := exec.Command("gm", "convert")
	return convertCmdConvertImage(cmd, imageData, width, height)
}
