package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ExchangeRate struct {
	UpdatedAt time.Time
	From      string
	To        string
}

func main() {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(
			"mongodb://localhost:27117/?replicaSet=TestMongoReplicaSet&connect=direct&readPreference=secondary",
		),
		options.Client().SetConnectTimeout(5*time.Second),
	)
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
		log.Fatalf("Insert error: %v", err)
	}
	log.Printf("Insert: %v", insert)

	options := options.Find()
	options.SetLimit(100)

	var rates []ExchangeRate
	cursor, err := collection.Find(context.TODO(), bson.M{}, options)
	if nil != err {
		log.Fatalf("Select error: %v", err)
	}
	err = cursor.All(context.TODO(), &rates)
	if nil != err {
		log.Fatalf("Fetch error: %v", err)
	}

	for _, r := range rates {
		log.Printf("%v", r)
	}
}
