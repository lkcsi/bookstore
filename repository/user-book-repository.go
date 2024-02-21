package repository

import "github.com/lkcsi/bookstore/entity"

type UserBookRepository interface {
	FindAll() ([]entity.UserBook, error)
	Find(string, string) (*entity.UserBook, error)
	Save(string, string) error
	Delete(string, string) error
	FindAllByUsername(string) ([]entity.UserBook, error)
}
