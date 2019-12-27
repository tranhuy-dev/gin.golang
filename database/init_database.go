package database

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbName = "GenBook"
)

var Client *mongo.Client
var DatabaseName string
var UrlDatabase string

func init() {
	SetupEnvironment()
	client, err := mongo.NewClient(options.Client().ApplyURI(UrlDatabase))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	Client = client
}

func SetupEnvironment() {
	// err := godotenv.Load()
	// if err != nil {
	//   log.Fatal("Error loading .env file")
	// }
	DatabaseName = "GenBook"
	UrlDatabase = "mongodb://localhost:27017/GenBook"
}