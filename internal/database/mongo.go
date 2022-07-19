package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	dbUrl := os.Getenv("DATABASE_URL")
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal("can't connect to the database")
	}
	return client
}
