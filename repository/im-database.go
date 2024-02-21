package repository

import (
	"github.com/lkcsi/bookstore/custerror"
	"github.com/lkcsi/bookstore/entity"
)

type ImDatabase struct {
	Books     []entity.Book
	UserBooks []entity.UserBook
}

func NewImDatabase() *ImDatabase {
	books := make([]entity.Book, 0)
	books = append(books, entity.Book{Id: "1", Title: "Title_1", Author: "Author_1", Quantity: quantity(5)})
	books = append(books, entity.Book{Id: "2", Title: "Title_2", Author: "Author_2", Quantity: quantity(5)})
	books = append(books, entity.Book{Id: "3", Title: "Title_3", Author: "Author_3", Quantity: quantity(5)})
	books = append(books, entity.Book{Id: "4", Title: "Title_4", Author: "Author_4", Quantity: quantity(5)})

	userBooks := make([]entity.UserBook, 0)

	return &ImDatabase{Books: books, UserBooks: userBooks}
}

func (bs *ImDatabase) FindBookById(id string) (*entity.Book, error) {
	for i, book := range bs.Books {
		if book.Id == id {
			return &bs.Books[i], nil
		}
	}
	return nil, custerror.BookNotFoundError(id)
}

func (bs *ImDatabase) FindBooksByUsername(username string) ([]entity.UserBook, error) {
	books := make([]entity.UserBook, 0)
	for _, book := range bs.UserBooks {
		if book.Username == username {
			books = append(books, book)
		}
	}
	return books, nil
}

func (bs *ImDatabase) Find(username, bookId string) (int, error) {
	for i, book := range bs.UserBooks {
		if book.Username == username && book.BookId == bookId {
			return i, nil
		}
	}
	return 0, custerror.UserBookNotFoundError(username, bookId)
}
