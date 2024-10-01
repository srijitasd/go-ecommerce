package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoService struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func MongoInit() (*MongoService, error) {
	var errorObj error

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, errorObj := mongo.Connect(clientOptions)
	if errorObj != nil {
		return nil, errorObj
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = client.Ping(ctx, readpref.Primary())
	fmt.Println("MongoDB connected successfully!")

	db := client.Database(os.Getenv("MONGO_DB"))

	return &MongoService{
		Client: client,
		DB:     db,
	}, nil
}

func (ms *MongoService) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("MongoDB disconnected successfully!")

	return ms.Client.Disconnect(ctx)
}
