//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package algo

type NullAlgo struct{}

func (a *NullAlgo) AlgoId() string {
	return "null"
}

func (a *NullAlgo) Process(args Args) (err error) {
	return
}
