package service

import (
	"github.com/lkcsi/bookstore/custerror"
	"github.com/lkcsi/bookstore/entity"
	"github.com/lkcsi/bookstore/repository"
)

type UserBookService interface {
	Checkout(string, string) (*entity.Book, error)
	FindAll() ([]entity.UserBook, error)
	Find(string, string) (*entity.UserBook, error)
	FindAllByUsername(string) ([]entity.UserBook, error)
	Return(string, string) error
}

type userBookService struct {
	userBookRepository repository.UserBookRepository
	bookRepository     repository.BookRepository
}

func NewUserBookService(ub *repository.UserBookRepository, b *repository.BookRepository) UserBookService {
	bs := userBookService{*ub, *b}
	return &bs
}

func (bs *userBookService) Checkout(username, id string) (*entity.Book, error) {
	book, err := bs.bookRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	if *book.Quantity < 1 {
		return nil, custerror.BookOutOfStockError(id)
	}

	_, err = bs.userBookRepository.Find(username, id)
	if err == nil {
		return nil, custerror.AlreadyCheckedError(username, id)
	}

	err = bs.userBookRepository.Save(username, id)
	if err != nil {
		return nil, err
	}

	*book.Quantity -= 1
	err = bs.bookRepository.Update(id, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (bs *userBookService) FindAll() ([]entity.UserBook, error) {
	return bs.userBookRepository.FindAll()
}

func (bs *userBookService) Find(username, bookId string) (*entity.UserBook, error) {
	return bs.userBookRepository.Find(username, bookId)
}

func (bs *userBookService) FindAllByUsername(username string) ([]entity.UserBook, error) {
	books, err := bs.userBookRepository.FindAllByUsername(username)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (bs *userBookService) Return(username, id string) error {
	book, err := bs.bookRepository.FindById(id)
	if err != nil {
		return custerror.BookNotFoundError(id)
	}

	*book.Quantity += 1

	err = bs.bookRepository.Update(id, book)
	if err != nil {
		return err
	}

	err = bs.userBookRepository.Delete(username, id)
	if err != nil {
		return err
	}

	return nil
}
