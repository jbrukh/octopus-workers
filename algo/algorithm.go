//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package algo

type AlgoSpec struct {
	resourceIds []string          `json: "resource_ids"`
	algoId      string            `json: "algo_id"`
	args        map[string]string `json: "args"`
}

func (a *AlgoSpec) ResourceIds() []string {
	return a.resourceIds
}

func (a *AlgoSpec) AlgoId() string {
	return a.algoId
}

func (a *AlgoSpec) Args() map[string]string {
	return a.args
}

// An algorithm
type Algorithm interface {
	AlgoId() string
	Process(resourceIds []string, args map[string]string)
}
