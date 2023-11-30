package mongo

import (
	"context"
	"fmt"
	"github.com/Minsoo-Shin/go-boilerplate/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func New(cfg config.Config) *mongo.Client {
	uri := fmt.Sprintf("mongodb+srv://%v:%v@%v/?retryWrites=true&w=majority",
		cfg.Mongo.User,
		cfg.Mongo.Password,
		cfg.Mongo.Host,
	)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetMinPoolSize(uint64(cfg.Mongo.Options.MinConnections)).
		SetMaxPoolSize(uint64(cfg.Mongo.Options.MaxConnections)).
		SetMaxConnIdleTime(10 * time.Minute)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
