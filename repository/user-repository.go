package repository

import "github.com/lkcsi/bookstore/entity"

type UserRepository interface {
	FindByUsername(string) (*entity.User, error)
	Save(*entity.User) error
}
