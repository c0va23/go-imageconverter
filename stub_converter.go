// +build !cmd
// +build !im
// +build !gm
// +build !vips

package main

import (
	"errors"
)

var err = errors.New("Build with tag 'im', 'gm', 'cmd im', 'cmd gm'")

const converterName = "stub"

func converterInitialize() {
	panic(err)
}

func converterTerminate() {}

func converter(imageData []byte, width, height uint) ([]byte, error) {
	return nil, err
}
