package auth

type User struct {
	ID          int64  `bson:"_id,omitempty"`
	PhoneNumber string `bson:"phone_number,omitempty"`
}
