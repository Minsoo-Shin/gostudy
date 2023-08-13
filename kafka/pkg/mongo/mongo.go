package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func New() (database *mongo.Database, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Println("database connection error", err)
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("Successfully connected and pinged.")

	return client.Database("test"), nil
}
