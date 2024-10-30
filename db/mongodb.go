package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	LinksDB     *mongo.Collection
)

func SetupMongoDB(uri string) error {
	var err error
	MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	LinksDB = MongoClient.Database("link-shortener").Collection("links")
	return nil
}
