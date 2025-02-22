package database

import "github.com/UnLess24/coin/client/internal/models/user"

type FakeDB struct{}

func NewFake() *FakeDB {
	return &FakeDB{}
}

func (f *FakeDB) FindUserByEmail(email string) (user.User, error) {
	return user.User{}, nil
}

func (f *FakeDB) CreateUser(user user.User) error {
	return nil
}
