package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ExchangeRate struct {
	UpdatedAt time.Time
	From      string
	To        string
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/sample"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("sample").Collection("exchangerate")
	insert, err := collection.InsertOne(context.TODO(), &ExchangeRate{UpdatedAt: time.Now(), From: "gold", To: "bitcoin"})
	if nil != err {
		log.Fatal(err)
	}

	log.Printf("Insert: %v", insert)
}
