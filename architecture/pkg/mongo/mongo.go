package mongo

import (
	"context"
	"fmt"
	"ggurugi/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	SetOperator  = "$set"
	PushOperator = "$push"
	InOperator   = "$in"
	NinOperator  = "$nin"
	AndOperator  = "$and"
	OrOperator   = "$or"
	GtOperator   = "$gt"
	GteOperator  = "$gte"
	LtOperator   = "$lt"
	LteOperator  = "$lte"
	EqOperator   = "$eq"
	NeOperator   = "$ne"
)

func New(cfg config.Config) *mongo.Client {
	uri := fmt.Sprintf("mongodb+srv://%v:%v@%v/?retryWrites=true&w=majority",
		cfg.Mongo.User,
		cfg.Mongo.Password,
		cfg.Mongo.Host,
	)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions).
		SetMinPoolSize(uint64(cfg.Mongo.Options.MinConnections)).
		SetMaxPoolSize(uint64(cfg.Mongo.Options.MaxConnections)).
		SetMaxConnIdleTime(10 * time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
