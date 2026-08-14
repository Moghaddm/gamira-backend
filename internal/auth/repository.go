package auth

type UserRepository interface {
	Create(user *User) (int64, error)
}
