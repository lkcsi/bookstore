package repository

import (
	"github.com/lkcsi/bookstore/entity"
)

type imUserBookRepository struct {
	db *ImDatabase
}

func ImUserBookRepository(db *ImDatabase) UserBookRepository {
	return &imUserBookRepository{db}
}

func (bs *imUserBookRepository) Save(username, bookId string) error {
	bs.db.UserBooks = append(bs.db.UserBooks, entity.UserBook{Username: username, BookId: bookId})
	return nil
}

func (bs *imUserBookRepository) Delete(username, bookId string) error {
	idx, err := bs.db.Find(username, bookId)
	if err != nil {
		return err
	}
	bs.db.UserBooks = append(bs.db.UserBooks[:idx], bs.db.UserBooks[idx+1:]...)
	return nil
}

func (bs *imUserBookRepository) Find(username, bookId string) (*entity.UserBook, error) {
	idx, err := bs.db.Find(username, bookId)
	if err != nil {
		return nil, err
	}
	return &bs.db.UserBooks[idx], nil
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
