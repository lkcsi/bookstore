package repository

import (
	"github.com/lkcsi/bookstore/custerror"
	"github.com/lkcsi/bookstore/entity"
)

type imBookRepository struct {
	db *ImDatabase
}

// DeleteAll implements BookRepository.
func (r *imBookRepository) DeleteAll() error {
	r.db.Books = make([]entity.Book, 0)
	return nil
}

func ImBookRepository(db *ImDatabase) BookRepository {
	return &imBookRepository{db: db}
}

func quantity(value int) *int {
	result := new(int)
	*result = value
	return result
}

func (bs *imBookRepository) Save(book *entity.Book) error {
	bs.db.Books = append(bs.db.Books, *book)
	return nil
}

func (bs *imBookRepository) Update(id string, updated *entity.Book) error {
	book, err := bs.db.FindBookById(id)
	if err != nil {
		return err
	}
	book.Author = updated.Author
	book.Quantity = updated.Quantity
	book.Title = updated.Title

	return nil
}

func (bs *imBookRepository) FindAll() ([]entity.Book, error) {
	return bs.db.Books, nil
}

func (bs *imBookRepository) FindById(id string) (*entity.Book, error) {
	return bs.db.FindBookById(id)
}

func (bs *imBookRepository) DeleteById(id string) error {
	index, err := bs.findBookIndex(id)
	if err != nil {
		return err
	}
	bs.db.Books = append(bs.db.Books[:index], bs.db.Books[index+1:]...)
	return nil
}

func (bs *imBookRepository) findBookIndex(id string) (int, error) {
	for i, book := range bs.db.Books {
		if book.Id == id {
			return i, nil
		}
	}
	return 0, custerror.BookNotFoundError(id)

}
