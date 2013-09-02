//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package main

import (
	"fmt"
	"github.com/jbrukh/octopus-workers/algo"
	"github.com/jonnii/go-workers"
	"log"
	"os"
)

// ----------------------------------------------------------------- //
// Constants
// ----------------------------------------------------------------- //

const (
	RedisServer      = "localhost:6379"
	RedisConnections = 2
	ProcessId        = "1"
	ProcessingQueue  = "sentipus-queue"
	Concurrency      = 2
	StatsServerPort  = 8888
)

func main() {
	redisServer := os.Getenv("REDIS_PROVIDER")
	if redisServer == "" {
		redisServer = RedisServer
	}

	password := os.Getenv("REDIS_PASSWORD")

	log.Printf("Starting up")
	log.Printf("SERVER: %s", redisServer)
	log.Printf("PASSWORD: %s", password)

	workers.Configure(map[string]string{
		"server":   redisServer,
		"pool":     fmt.Sprintf("%s", RedisConnections),
		"process":  ProcessId,
		"password": password,
	})

	workers.Middleware.Append(&workers.MiddlewareLogging{})
	workers.Middleware.Append(&MiddlewareResponder{})

	workers.Process(ProcessingQueue, RunOctopusWorker, Concurrency)
	workers.Run()
}

func RunOctopusWorker(args *workers.Args) {
	log.Printf("received message: %s", args.ToJson())

	a1 := args.GetIndex(0)
	if a1 == nil {
		log.Printf("Couldn't extract the first item of the args array")
		return
	}

	algoId, err := a1.Get("algo_id").String()
	if err != nil {
		log.Print("Couldn't find a key in the worker request with name 'algo_id'")
		return
	}

	a := a1.Get("args")
	if a == nil {
		log.Printf("Couldn't find a key in the worker request with the name 'args'")
		return
	}

	//algorithm, err := algo.Provide(algoId)
	algorithm, err := algo.Provide("null")
	if err != nil {
		log.Printf("Couldn't find algorithm with name: %s", algoId)
		return
	}

	log.Printf("Starting algorithm: %s", algoId)

	err = algorithm.Process(algo.Args(a))
	if err != nil {
		log.Printf("An error occured while processing an algorithm, %v", err)
	}
}
