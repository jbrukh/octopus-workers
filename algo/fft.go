//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package algo

import (
	"fmt"
	"github.com/jbrukh/goavatar/formats"
)

type FftAlgo struct{}

func (a *FftAlgo) AlgoId() string {
	return "fft"
}

func (a *FftAlgo) Process(args Args) (err error) {
	r, _ := formats.NewOBFReader(nil)
	fmt.Println(r)
	return
}
