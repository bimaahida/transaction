package database

import (
	"context"
	"time"
	"transaction/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(conf config.Database) (*mongo.Database, *context.Context) {
	connection := options.Client().ApplyURI(conf.Server)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, connection)

	if err != nil {
		panic(err)
	}

	return client.Database(conf.Database), &ctx
}
