package main

import (
	"github.com/jrallison/go-workers"
	"testing"
)

func TestRunOctopusWorker(t *testing.T) {
	msg, _ := workers.NewMsg(`{"args":[{"algo_id":"null","args":{"input_file":"input-file-name"}}]}`)
	RunOctopusWorker(msg.Args())
}
