//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package algo

import (
	"fmt"
)

// Available algorithms.
var algos = map[string]Algo{
	"fft":     new(FftAlgo),
	"average": new(AverageAlgo),
	"null":    new(NullAlgo),
}

// Provide an algorithm by id.
func Provide(algoId string) (algo Algo, err error) {
	if algo, ok := algos[algoId]; !ok {
		return nil, fmt.Errorf("unknown algo: %s", algoId)
	} else {
		return algo, nil
	}
	return
}
