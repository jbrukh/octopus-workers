package main

import (
	"encoding/json"
	"github.com/jonnii/go-workers"
	"log"
	"time"
)

type Response struct {
	Jid        string        `json:"jid"`
	ClassName  string        `json:"class"`
	EnqueuedAt int64         `json:"enqueued_at"`
	Args       []interface{} `json:"args"`
}

type MiddlewareResponder struct{}

func (l *MiddlewareResponder) Call(queue string, message *workers.Msg, next func()) {
	defer func() {
		log.Printf("Responding to message: %s", message.Jid())

		conn := workers.Config.GetConn()
		defer conn.Close()

		_, err := conn.Do("sadd", "queues", "default")
		if err != nil {
			log.Printf("Could not create default queue %v", err)
			return
		}

		message := Response{}
		message.Jid = "1234"
		message.ClassName = "GoResponseWorker"
		message.EnqueuedAt = time.Now().Unix()
		message.Args = make([]interface{}, 0)

		response, err := json.Marshal(message)
		if err != nil {
			log.Printf("Could nto marshal response to json, %v", err)
		}

		conn.Do("lpush", "queue:default", response)
	}()

	next()
}
