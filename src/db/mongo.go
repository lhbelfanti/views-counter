package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDatabase struct {
	client *mongo.Client
	ctx    context.Context
}

// NewMongoDatabase creates a new *MongoDatabase
func NewMongoDatabase() *MongoDatabase {
	uri := os.Getenv("MONGO_URL")

	mongoDB := &MongoDatabase{}

	client, ctx, err := mongoDB.Connect(uri)
	if err != nil {
		panic(err)
	}

	mongoDB.client = client
	mongoDB.ctx = ctx

	return mongoDB
}

// GetCurrentCount returns the current view count
func (mongoDB *MongoDatabase) GetCurrentCount() int {
	// Select the correct database and collection
	collection := mongoDB.client.Database("prod").Collection("counter")

	// Define your filter
	filter := bson.D{}

	// Search for the item
	var result bson.M
	err := collection.FindOne(mongoDB.ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	// Access the value of the "views" key
	viewsRaw, ok := result["views"]
	if !ok {
		log.Fatal("Key 'views' not found")
	}

	viewsInt32, ok := viewsRaw.(int32)
	if !ok {
		log.Fatal("Unable to cast 'views' value to int")
	}

	return int(viewsInt32)
}

// UpdateCurrentCount updates the current view count and returns it
func (mongoDB *MongoDatabase) UpdateCurrentCount() int {
	// Select the correct database and collection
	collection := mongoDB.client.Database("prod").Collection("counter")

	currentCount := mongoDB.GetCurrentCount()
	currentCount++

	// Define your filter
	filter := bson.D{}
	update := bson.D{{"$set", bson.D{
		{"views", currentCount},
	}},
	}

	// Search for the item
	_, err := collection.UpdateOne(mongoDB.ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return currentCount
}

// Close method to close resources.
// This method closes mongoDB connection and cancel context.
func (mongoDB *MongoDatabase) Close() {
	// client provides a method to close a mongoDB connection.
	defer func() {
		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := mongoDB.client.Disconnect(mongoDB.ctx); err != nil {
			panic(err)
		}
	}()
}

// Connect returns mongo.Client, context.Context and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
func (mongoDB *MongoDatabase) Connect(uri string) (*mongo.Client, context.Context, error) {
	ctx := context.TODO()
	ops := options.Client().ApplyURI(uri)
	maxConnIdleTime := 60 * time.Second
	ops.MaxConnIdleTime = &maxConnIdleTime

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, ops)

	return client, ctx, err
}

// Ping is used to ping the mongoDB, return error if any
func (mongoDB *MongoDatabase) Ping() error {
	// mongo.Client has Ping to ping mongoDB, deadline of the Ping method will be determined by cxt
	// Ping method return error if any occurred, then the error can be handled.
	if err := mongoDB.client.Ping(mongoDB.ctx, readpref.Primary()); err != nil {
		return err
	}

	fmt.Println("connected successfully")
	return nil
}
