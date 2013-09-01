//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package algo

import (
	"github.com/bitly/go-simplejson"
)

// Args is a wrapper around simplejson
type Args simplejson.Json

// An algorithm
type Algorithm interface {
	AlgoId() string
	Process(args *Args) (err error)
}
