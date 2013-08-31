//
// Copyright (c) 2013 Jake Brukhman/East River Labs. All rights reserved.
//
package main

import (
	"fmt"
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

	workers.Process(ProcessingQueue, sentipusWorker, Concurrency)
	workers.Run()

}

func sentipusWorker(args *workers.Args) {
	log.Printf("received message: %v", args)
}
