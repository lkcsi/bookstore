package repository

import "github.com/lkcsi/bookstore/entity"

type UserBookRepository interface {
	FindAll() ([]entity.UserBook, error)
	Checkout(string, string) error
	FindAllByUsername(string) ([]entity.UserBook, error)
	Return(string, string) error
}
