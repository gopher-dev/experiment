package main

import (
	"log"

	"github.com/gocraft/work"
	"github.com/gopher-dev/experiment/gocraft-work/config"
)

// Make an enqueuer with a particular namespace
var enqueuer = work.NewEnqueuer("my_app_namespace", config.RedisPool)

func main() {
	// Enqueue a job named "send_email" with the specified parameters.
	_, err := enqueuer.Enqueue("send_email", work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4})
	if err != nil {
		log.Fatal(err)
	}
}
