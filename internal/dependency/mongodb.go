package dependency

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(config Config, logger Logger) (*mongo.Database, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDB.MongoUri))
	if err != nil {
		logger.Errorf("Error connecting with MongoDB", err)
	}
	mongodb := client.Database(config.MongoDB.MongoName)

	logger.Infof("Successfully Connect to MongoDB", nil)

	return mongodb, nil
}
