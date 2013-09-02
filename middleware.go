package main

import (
	"encoding/json"
	"github.com/jonnii/go-workers"
	"log"
	"math/rand"
	"time"
)

type Response struct {
	Jid        string   `json:"jid"`
	ClassName  string   `json:"class"`
	EnqueuedAt int64    `json:"enqueued_at"`
	Args       []string `json:"args"`
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

		a1 := message.Args().GetIndex(0)
		if a1 == nil {
			log.Printf("Couldn't extract the first item of the args array")
			return
		}

		analysisId, err := a1.Get("analysis_id").String()
		if err != nil {
			log.Print("Couldn't find a key in the worker message with name 'analysis_id'")
			log.Print("ERR: %v", err)
			return
		}

		message := Response{
			randomString(20),
			"GoResponseWorker",
			time.Now().Unix(),
			[]string{analysisId},
		}

		response, err := json.Marshal(message)
		if err != nil {
			log.Printf("Could nto marshal response to json, %v", err)
		}

		conn.Do("lpush", "queue:default", response)
	}()

	next()
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
