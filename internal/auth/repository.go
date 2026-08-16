package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, phoneNumber string) (string, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*User, error)
	GetById(ctx context.Context, id string) (*User, error)
}

type repository struct {
	users *mongo.Collection
}

func NewRepository(users *mongo.Collection) Repository {
	return &repository{users: users}
}

func (repo *repository) Create(ctx context.Context, phoneNumber string) (string, error) {
	result, err := repo.users.InsertOne(ctx, &User{PhoneNumber: phoneNumber})
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo *repository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*User, error) {
	user := &User{}
	err := repo.users.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *repository) GetById(ctx context.Context, id string) (*User, error) {
	user := &User{}
	err := repo.users.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
