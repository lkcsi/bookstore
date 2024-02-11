package repository

import "github.com/lkcsi/bookstore/entity"

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	Save(*entity.Book) error
	Update(string, *entity.Book) error
	FindById(string) (*entity.Book, error)
	DeleteById(string) error
	DeleteAll() error
}
