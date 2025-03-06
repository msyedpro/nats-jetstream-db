package main

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	//Setup Jet Stream Context
	ctx := context.Background()
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	//Create Publisher Jet Stream
	_, err = js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:        "items",
		Description: "Messages for items",
		Subjects: []string{
			"items.>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Publish messages
	js.Publish(ctx, "items.1", nil)
	log.Printf("Publish message %d", 1)
	js.Publish(ctx, "items.2", nil)
	log.Printf("Publish message %d", 2)
	js.Publish(ctx, "items.3", nil)
	log.Printf("Publish message %d", 3)
	js.Publish(ctx, "items.1", nil)
	log.Printf("Publish message %d", 4)
	js.Publish(ctx, "items.2", nil)
	log.Printf("Publish message %d", 5)
	js.Publish(ctx, "items.3", nil)
	log.Printf("Publish message %d", 6)
	js.Publish(ctx, "items.1", nil)
	log.Printf("Publish message %d", 7)
	js.Publish(ctx, "items.2", nil)
	log.Printf("Publish message %d", 8)
	js.Publish(ctx, "items.3", nil)
	log.Printf("Publish message %d", 9)
	js.Publish(ctx, "items.1", nil)
	log.Printf("Publish message %d", 10)
	js.Publish(ctx, "items.2", nil)
}
