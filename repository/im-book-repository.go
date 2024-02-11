package repository

import (
	"github.com/lkcsi/bookstore/custerror"
	"github.com/lkcsi/bookstore/entity"
)

type imBookRepository struct {
	books []entity.Book
}

// DeleteAll implements BookRepository.
func (r *imBookRepository) DeleteAll() error {
	r.books = make([]entity.Book, 0)
	return nil
}

func NewImBookRepository() BookRepository {
	books := make([]entity.Book, 0)
	return &imBookRepository{books: books}
}

func (bs *imBookRepository) Save(book *entity.Book) error {
	bs.books = append(bs.books, *book)
	return nil
}

func (bs *imBookRepository) Update(id string, updated *entity.Book) error {
	book, err := bs.findBookById(id)
	if err != nil {
		return err
	}
	book.Author = updated.Author
	book.Quantity = updated.Quantity
	book.Title = updated.Title

	return nil
}

func (bs *imBookRepository) FindAll() ([]entity.Book, error) {
	return bs.books, nil
}

func (bs *imBookRepository) FindById(id string) (*entity.Book, error) {
	return bs.findBookById(id)
}

func (bs *imBookRepository) DeleteById(id string) error {
	index, err := bs.findBookIndex(id)
	if err != nil {
		return err
	}
	bs.books = append(bs.books[:index], bs.books[index+1:]...)
	return nil
}

func (bs *imBookRepository) findBookIndex(id string) (int, error) {
	for i, book := range bs.books {
		if book.Id == id {
			return i, nil
		}
	}
	return 0, custerror.BookNotFoundError(id)

}

func (bs *imBookRepository) findBookById(id string) (*entity.Book, error) {
	for i, book := range bs.books {
		if book.Id == id {
			return &bs.books[i], nil
		}
	}
	return nil, custerror.BookNotFoundError(id)
}
