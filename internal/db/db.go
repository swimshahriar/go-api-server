package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var mongoOnce sync.Once

var clinetInstanceError error

type Collection string

const (
	ProductsCollection Collection = "products"
)

const (
	uri      = "mongodb://localhost:27017"
	Database = "products-api"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client

		clinetInstanceError = err
	})
	return clientInstance, clinetInstanceError
}
