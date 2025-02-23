package user

import (
	"time"

	"github.com/UnLess24/coin/client/internal/dto"
)

type User struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

func FromRegisterRequest(req dto.RegisterRequest) User {
	return User{
		Email:    req.Email,
		Password: req.Password,
	}
}
