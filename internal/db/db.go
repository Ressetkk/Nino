package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var client *mongo.Client
var db *mongo.Database

const Name = "nino_db"

func InitDBConnection() error {
	log.Info("Waiting for DB connection...")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return err
	}
	if err := c.Ping(ctx, nil); err != nil {
		return err
	}
	client = c
	db = client.Database(Name)
	return nil
}

func GetCollection(collection string) *mongo.Collection {
	return db.Collection(collection)
}
