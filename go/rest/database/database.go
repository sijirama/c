package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// function to start a MongoDB connection
// return error and database client

var DBclient *mongo.Client

func StartMongoDb() (error, *mongo.Client) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    DBclient = client
	if err != nil {
		log.Panic("Error connecting to database")
		return errors.New("Error connecting to database"), nil
	}
	return nil, client
}

func CloseMongoDb(client *mongo.Client) error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return errors.New(fmt.Sprintf("Error disconnecting from database: %v", err))
	}
	return nil
}

func GetCollection(client mongo.Client, name string) *mongo.Collection {
	coll := client.Database("test").Collection(name)
	return coll
}
