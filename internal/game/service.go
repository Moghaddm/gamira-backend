package game

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	Repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{Repository: repo}
}

func (s Service) SeedAll(ctx context.Context) error {
	games := []Game{
		{ID: primitive.NewObjectID(), Name: "Game 1", IconUrl: ""},
		{ID: primitive.NewObjectID(), Name: "Game 2", IconUrl: ""},
		{ID: primitive.NewObjectID(), Name: "Game 3", IconUrl: ""},
	}

	for _, game := range games {
		e, err := s.Repository.Exists(ctx, game.ID)
		if err != nil {
			return err
		}

		if !e {
			err := s.Repository.Create(ctx, &game)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
