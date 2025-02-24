package database

import (
	"github.com/UnLess24/coin/client/internal/models/user"
)

var (
	ErrUserAlreadyExists = "user already exists"
)

type DB interface {
	FindUserByEmail(email, pass string) (user.User, error)
	CreateUser(user user.User) error
}
