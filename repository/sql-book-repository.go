package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lkcsi/bookstore/custerror"
	"github.com/lkcsi/bookstore/entity"
)

type sqlBookRepository struct {
	db *sql.DB
}

func SqlBookRepository() BookRepository {
	pwd := os.Getenv("BOOKS_DB_PASSWORD")
	port := os.Getenv("BOOKS_DB_PORT")
	host := os.Getenv("BOOKS_DB_HOST")
	conn := fmt.Sprintf("root:%s@tcp(%s:%s)/book_db", pwd, host, port)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println("Unable to connect database")
	}
	return &sqlBookRepository{db}
}
func (r *sqlBookRepository) DeleteAll() error {
	_, err := r.db.Exec("DELETE FROM books")
	if err != nil {
		return err
	}
	return nil
}

func (repo *sqlBookRepository) FindAll() ([]entity.Book, error) {
	res, err := repo.db.Query("SELECT id, title, author, quantity FROM books")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	result := make([]entity.Book, 0)
	for res.Next() {
		var book entity.Book
		err := res.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			return nil, err
		}
		result = append(result, book)
	}
	return result, nil
}

func (repo *sqlBookRepository) DeleteById(id string) error {
	res, err := repo.db.Exec("DELETE FROM books WHERE id=?", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return custerror.BookNotFoundError(id)
	}
	return nil
}

func (repo *sqlBookRepository) FindById(id string) (*entity.Book, error) {
	var book entity.Book
	row := repo.db.QueryRow("SELECT id, title, author, quantity FROM books WHERE id=?", id)
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (repo *sqlBookRepository) Save(b *entity.Book) error {
	_, err := repo.db.Exec("INSERT INTO books (id,title,author,quantity) VALUES(?, ?, ?, ?)", b.Id, b.Title, b.Author, b.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (repo *sqlBookRepository) Update(id string, b *entity.Book) error {
	_, err := repo.db.Exec("UPDATE books SET title=?, author=?, quantity=? WHERE id=?", b.Title, b.Author, b.Quantity, id)
	if err != nil {
		return err
	}
	return nil
}
