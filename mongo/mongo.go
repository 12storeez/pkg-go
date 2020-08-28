package mongo

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhs/loggr"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongo(url string, dbName string) (*Mongo, error) {
	log := loggr.WithContext(context.Background())
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, errors.Errorf("mongo.NewMongo.NewClient: %v", err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, errors.Errorf("mongo.NewMongo.Connect: %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	log.Info("connected to mongodb")
	return &Mongo{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}
