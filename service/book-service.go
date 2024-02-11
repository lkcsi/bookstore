package service

import (
	"github.com/google/uuid"
	"github.com/lkcsi/bookstore/custerror"
	"github.com/lkcsi/bookstore/entity"
	"github.com/lkcsi/bookstore/repository"
)

type BookService interface {
	Save(entity.Book) (*entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindById(string) (*entity.Book, error)
	DeleteById(string) error
	Checkout(string) (*entity.Book, error)
	DeleteAll() error
}

type bookService struct {
	bookRepository repository.BookRepository
}

func (b *bookService) DeleteAll() error {
	return b.bookRepository.DeleteAll()
}

func InMemoryBookService() BookService {
	repo := repository.NewImBookRepository()
	bs := bookService{repo}
	return &bs
}

func NewSqlBookService() BookService {
	repo := repository.NewSqlBookRepository()
	return &bookService{repo}
}

func (bs *bookService) Save(book entity.Book) (*entity.Book, error) {
	book.Id = uuid.NewString()
	if err := bs.bookRepository.Save(&book); err != nil {
		return nil, err
	}
	return &book, nil
}

func (bs *bookService) FindAll() ([]entity.Book, error) {
	return bs.bookRepository.FindAll()
}

func (bs *bookService) FindById(id string) (*entity.Book, error) {
	book, err := bs.bookRepository.FindById(id)
	if err != nil {
		return nil, custerror.BookNotFoundError(id)
	}
	return book, nil
}

func (bs *bookService) DeleteById(id string) error {
	if err := bs.bookRepository.DeleteById(id); err != nil {
		return err
	}
	return nil
}

func (bs *bookService) Checkout(id string) (*entity.Book, error) {
	book, err := bs.bookRepository.FindById(id)
	if err != nil {
		return nil, custerror.BookNotFoundError(id)
	}
	if *book.Quantity == 0 {
		return nil, custerror.BookOutOfStockError(id)
	}
	*book.Quantity -= 1
	if err := bs.bookRepository.Update(id, book); err != nil {
		return nil, err
	}
	return book, nil
}
