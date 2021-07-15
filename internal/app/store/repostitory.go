package store

import "gorest2/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
	Find(int) (*model.User, error)
}
