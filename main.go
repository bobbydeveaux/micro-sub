package main

import (
	"encoding/json"

	"github.com/nats-io/go-nats"
	"log"
	"os"

	"time"
)

type person struct {
	Id       int64
	Name     string
	HostName string
}

func main() {
	//NATS_HOST = nats://localhost:4222
	nc, _ := nats.Connect(os.Getenv("NATS_HOST"))
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()

	var p person
	ec.QueueSubscribe("newuser", "job_workers", func(msg *nats.Msg) {
		log.Printf("Received a person: %s\n", msg.Data)
		err := json.Unmarshal(msg.Data, &p)
		if err != nil {
			log.Println(err.Error())
		}
		// @TODO create user in database
		p.Id = int64(time.Now().UnixNano())
		p.HostName, err = os.Hostname()
		if err != nil {
			log.Println(err.Error())
		}
		ec.Publish(msg.Reply, p)
	})
	select {}
}
