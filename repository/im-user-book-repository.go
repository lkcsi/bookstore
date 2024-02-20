package repository

import (
	"github.com/lkcsi/bookstore/custerror"
	"github.com/lkcsi/bookstore/entity"
)

type imUserBookRepository struct {
	db *ImDatabase
}

func ImUserBookRepository(db *ImDatabase) UserBookRepository {
	return &imUserBookRepository{db}
}

func (bs *imUserBookRepository) Checkout(username, id string) error {
	_, err := bs.db.Find(username, id)
	if err == nil {
		return custerror.AlreadyCheckedError(id, username)
	}
	book, err := bs.db.FindBookById(id)
	if err != nil {
		return err
	}

	*book.Quantity -= 1
	bs.db.UserBooks = append(bs.db.UserBooks, entity.UserBook{Username: username, BookId: id})
	return nil
}

func (bs *imUserBookRepository) FindAll() ([]entity.UserBook, error) {
	userBooks := make([]entity.UserBook, len(bs.db.UserBooks))
	for idx, userBook := range bs.db.UserBooks {
		book, _ := bs.db.FindBookById(userBook.BookId)
		userBook.Title = book.Title
		userBook.Author = book.Author
		userBooks[idx] = userBook
	}
	return userBooks, nil
}

func (bs *imUserBookRepository) FindAllByUsername(username string) ([]entity.UserBook, error) {
	resp, _ := bs.db.FindBooksByUsername(username)
	userBooks := make([]entity.UserBook, len(resp))
	for idx, userBook := range resp {
		book, _ := bs.db.FindBookById(userBook.BookId)
		userBook.Title = book.Title
		userBook.Author = book.Author
		userBooks[idx] = userBook
	}
	return userBooks, nil
}

func (bs *imUserBookRepository) Return(username, bookId string) error {
	index, err := bs.db.Find(username, bookId)
	if err != nil {
		return err
	}
	book, err := bs.db.FindBookById(bookId)
	if err != nil {
		return err
	}
	*book.Quantity += 1
	bs.db.UserBooks = append(bs.db.UserBooks[:index], bs.db.UserBooks[index+1:]...)
	return nil
}
