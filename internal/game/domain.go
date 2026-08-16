package game

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	IconUrl string             `bson:"icon_url"`
}
