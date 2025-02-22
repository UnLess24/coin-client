package database

import (
	"fmt"
	"sync"

	"github.com/UnLess24/coin/client/internal/models/user"
)

type FakeDB struct {
	mu    sync.RWMutex
	store map[string]user.User
}

func NewFake() *FakeDB {
	return &FakeDB{
		store: make(map[string]user.User),
	}
}

func (f *FakeDB) FindUserByEmail(email, pass string) (user.User, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	u, ok := f.store[email]
	if !ok {
		return user.User{}, fmt.Errorf("user or password is incorrect")
	}

	if u.Password != pass {
		return user.User{}, fmt.Errorf("user or password is incorrect")
	}

	return u, nil
}

func (f *FakeDB) CreateUser(user user.User) error {
	if user.Email == "" || user.Password == "" {
		return fmt.Errorf("email or password is incorrect")
	}

	f.mu.RLock()
	_, ok := f.store[user.Email]
	f.mu.RUnlock()
	if ok {
		return fmt.Errorf("user already exists")
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	f.store[user.Email] = user

	return nil
}
