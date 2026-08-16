package game

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, game *Game) error
	Exists(ctx context.Context, id primitive.ObjectID) (bool, error)
}

type repository struct {
	games *mongo.Collection
}

func NewRepository(games *mongo.Collection) Repository {
	return &repository{games: games}
}

func (repo *repository) Create(ctx context.Context, game *Game) error {
	_, err := repo.games.InsertOne(ctx, game)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) Exists(ctx context.Context, id primitive.ObjectID) (bool, error) {
	err := repo.games.FindOne(ctx, bson.M{"_id": id}).Decode(&Game{})
	if err != nil {
		return false, err
	}
	return true, nil
}
