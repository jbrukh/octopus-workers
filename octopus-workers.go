//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package main

import (
	"fmt"
	"github.com/jbrukh/octopus-workers/algo"
	"github.com/jrallison/go-workers"
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

	workers.Process(ProcessingQueue, sentipusWorker, Concurrency)
	workers.Run()
}

func sentipusWorker(args *workers.Args) {
	log.Printf("received message: %v", args)

	algoId, err := args.Get("algo_id").String()
	if err != nil {
		log.Print("Couldn't find a key in the worker request with name 'algo_id'")
		return
	}

	a := args.Get("args")
	if a != nil {
		log.Printf("Couldn't find a key in the worker request with the name 'args'")
		return
	}

	algorithm, err := algo.Provide(algoId)
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
