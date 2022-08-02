package store

import "github.com/Dennikoff/UserTagApi/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
}
