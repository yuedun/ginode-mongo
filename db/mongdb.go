package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var Mongodb *mongo.Database

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://10.168.160.14:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}
	Mongodb = client.Database("website")
}
