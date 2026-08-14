package game

type Game struct {
	ID      int64  `bson:"_id,omitempty"`
	Name    string `bson:"name"`
	IconUrl string `bson:"icon_url"`
}
