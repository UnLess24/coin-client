package database

import (
	"github.com/UnLess24/coin/client/internal/models/user"
)

type DB interface {
	FindUserByEmail(email, pass string) (user.User, error)
	CreateUser(user user.User) error
}
