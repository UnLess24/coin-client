package database

import (
	"context"
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

func (f *FakeDB) FindUserByEmail(ctx context.Context, email, pass string) (user.User, error) {
	err := proceedWithCheckContext(ctx, func() error {
		return nil
	})
	if err != nil {
		return user.User{}, err
	}

	f.mu.RLock()
	defer f.mu.RUnlock()

	u, ok := f.store[email]
	if !ok {
		return user.User{}, ErrEmailOrPasswordIsIncorrect
	}

	if u.Password != pass {
		return user.User{}, ErrEmailOrPasswordIsIncorrect
	}

	return u, nil
}

func (f *FakeDB) CreateUser(ctx context.Context, user user.User) error {
	err := proceedWithCheckContext(ctx, func() error {
		if user.Email == "" || user.Password == "" {
			return ErrEmailOrPasswordIsIncorrect
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = proceedWithCheckContext(ctx, func() error {
		f.mu.RLock()
		_, ok := f.store[user.Email]
		f.mu.RUnlock()
		if ok {
			return ErrUserAlreadyExists
		}
		return nil
	})
	if err != nil {
		return err
	}

	return proceedWithCheckContext(ctx, func() error {
		f.mu.Lock()
		defer f.mu.Unlock()
		f.store[user.Email] = user

		return nil
	})
}

func (f *FakeDB) Close() error {
	return nil
}

func proceedWithCheckContext(ctx context.Context, fn func() error) error {
	select {
	case <-ctx.Done():
		return ErrContextIsCanceled
	default:
		return fn()
	}
}
