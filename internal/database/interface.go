package database

import (
	"context"
	"errors"

	"github.com/UnLess24/coin/client/internal/models/user"
)

var (
	ErrUserAlreadyExists          = errors.New("user already exists")
	ErrEmailOrPasswordIsIncorrect = errors.New("email or password is incorrect")
	ErrContextIsCanceled          = errors.New("context is canceled")
)

type DB interface {
	FindUserByEmail(ctx context.Context, email, pass string) (user.User, error)
	CreateUser(ctx context.Context, user user.User) error
}
