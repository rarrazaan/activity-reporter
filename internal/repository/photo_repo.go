package repository

import (
	"context"
	"mini-socmed/internal/cons"
	"mini-socmed/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	photoRepo struct {
		mongo *mongo.Database
	}
	PhotoRepo interface {
		AddPhoto(ctx context.Context, photo *model.Photo) (*mongo.InsertOneResult, error)
	}
)

func (pr *photoRepo) AddPhoto(ctx context.Context, photo *model.Photo) (*mongo.InsertOneResult, error) {
	res, err := pr.mongo.Collection(cons.MongoUserPostCollection).InsertOne(ctx, photo)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewPhotoRepo(mongo *mongo.Database) PhotoRepo {
	return &photoRepo{
		mongo: mongo,
	}
}
