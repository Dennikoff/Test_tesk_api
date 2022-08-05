package teststore

import (
	"errors"
	"github.com/Dennikoff/UserTagApi/internal/app/model"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(user *model.User) error {
	_, ok := r.users[user.Email]
	if ok {
		return errors.New("duplicate key")
	}

	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	r.users[user.Email] = user
	user.ID = uint(len(r.users))

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user, ok := r.users[email]
	if !ok {
		return nil, errors.New("not found")
	}
	return user, nil
}
