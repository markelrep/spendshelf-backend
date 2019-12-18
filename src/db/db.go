package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Database struct {
	MongoDB    *mongo.Database
	logger     *zap.SugaredLogger
	CancelFunc context.CancelFunc
}

// NewDatabase is create a new database connection
func NewDatabase(dbname, URI string) (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI), options.Client().SetMaxPoolSize(50))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*4999)
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	database := client.Database(dbname)

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Database{
		MongoDB: database,
		logger:  logger.Sugar(),
	}, nil
}