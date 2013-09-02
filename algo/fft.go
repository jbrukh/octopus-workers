//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package algo

import (
	"fmt"
	. "github.com/jbrukh/octopus-workers/resources"
	"io"
)

//
// Inputs:
//    resource_id: reso
//
type FftAlgo struct{}

func (a *FftAlgo) AlgoId() string {
	return "fft"
}

func (a *FftAlgo) Process(args Args) (err error) {
	resource, err := args.Get("resource").String()
	if err != nil {
		return fmt.Errorf("missing parameter: resource")
	}

	storage, err := args.Get("storage").String()
	if err != nil || (storage != "s3" && storage != "file") {
		return fmt.Errorf("missing parameter: storage=[s3|file]")
	}

	var r io.Reader
	switch storage {
	case "s3":
		if r, err = GetUrl(resource); err != nil {
			return err
		}
	case "file":
		if r, err = GetFile(resource); err != nil {
			return err
		}
	default:
		return fmt.Errorf("missing parameter: storage=[s3|file]")
	}
	fmt.Printf("%v", r)

	return
}
