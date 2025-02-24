package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/UnLess24/coin/client/config"
	"github.com/UnLess24/coin/client/internal/models/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type PGDB struct {
	db *sqlx.DB
}

func NewPGDB(cfg *config.Config) (*PGDB, error) {
	dsName := fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Database, cfg.DB.SslMode)
	db, err := sqlx.Connect(cfg.DB.Name, dsName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &PGDB{db: db}, nil
}

func (p *PGDB) FindUserByEmail(ctx context.Context, email, pass string) (user.User, error) {
	var u user.User
	err := p.db.GetContext(ctx, &u, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		slog.Error("failed to find user by email", "error", err)
		return user.User{}, ErrEmailOrPasswordIsIncorrect
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		slog.Error("failed to user password", "error", err)
		return user.User{}, ErrEmailOrPasswordIsIncorrect
	}

	return u, nil
}

func (p *PGDB) CreateUser(ctx context.Context, user user.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to generate hash user password", "error", err)
		return ErrEmailOrPasswordIsIncorrect
	}
	user.Password = string(hash)

	_, err = p.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	if err != nil {
		slog.Error("failed to insert user into db", "error", err)
		return ErrUserAlreadyExists
	}

	return nil
}

func (p *PGDB) Close() error {
	return p.db.Close()
}
