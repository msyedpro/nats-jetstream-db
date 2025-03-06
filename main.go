package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	ctx := context.Background()

	//Setup Jetstream with Context
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	stream, err := js.Stream(ctx, "items")
	if err != nil {
		log.Fatal(err)
	}

	//Setup Consumer
	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:    "items_consumer",
		Durable: "items_consumer",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Consume the data from Publisher with acknowledgment
	cctx, err := consumer.Consume(func(msg jetstream.Msg) {
		log.Printf("Received message: %s", string(msg.Subject()))
		msg.Ack()

		// add the cosumed message to database
		addItemToDatabase(string(msg.Subject()))
	})

	// fmt.Println("# Fetch messages")
	// consumer, _ = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
	// 	InactiveThreshold: 10 * time.Millisecond,
	// })

	// fetchResult, _ := consumer.Fetch(3, jetstream.FetchMaxWait(100*time.Millisecond))
	// for msg := range fetchResult.Messages() {
	// 	fmt.Printf("received %q\n", msg.Subject())
	//  addToDatabase(string(msg.Subject()))
	// 	msg.Ack()
	// }

	if err != nil {
		log.Fatal(err)
	}
	defer cctx.Stop()

	//Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

type Item struct {
	Name string
}

func addItemToDatabase(itemname string) {
	connStr := "postgres://dbuser:dbpass@localhost:5432/dbuser?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createItemTable(db)
	item := Item{itemname}
	insertItem(db, item)

	//var name string

	// query := `SELECT name FROM item`
	// err = db.QueryRow(query).Scan(&name)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Fatalf("No rows found with ID %d", pk)
	// 	}
	// 	log.Fatal(err)
	// }
	// fmt.Printf("AddedItemName: %s\n", name)
}

func createItemTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS item(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

func insertItem(db *sql.DB, item Item) int {
	query := `INSERT INTO item (name)
	VALUES ($1) RETURNING id`

	var pk int
	err := db.QueryRow(query, item.Name).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
