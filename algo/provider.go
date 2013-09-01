//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package algo

import (
	"fmt"
)

// Available algorithms.
var algorithms = map[string]Algorithm {
	"fft": new(FftAlgo),
	"average": new(AverageAlgo),
}

// Provide an algorithm by id.
func Provide(algoId string) (algo Algorithm, err error) {
	if algo, ok := algorithms[algoId]; !ok {
		return nil, fmt.Errorf("unknown algo: %s", algoId)
	} else {
		return algo, nil
	}
	return
}

