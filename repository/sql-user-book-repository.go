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

type sqlUserBookRepository struct {
	db *sql.DB
}

func SqlUserBookRepository() UserBookRepository {
	pwd := os.Getenv("BOOKS_DB_PASSWORD")
	port := os.Getenv("BOOKS_DB_PORT")
	host := os.Getenv("BOOKS_DB_HOST")
	conn := fmt.Sprintf("root:%s@tcp(%s:%s)/book_db", pwd, host, port)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println("Unable to connect database")
	}
	return &sqlUserBookRepository{db}
}

func (repo *sqlUserBookRepository) FindAll() ([]entity.UserBook, error) {
	res, err := repo.db.Query(`SELECT books.id, books.title, books.author, user_books.username 
	FROM user_books 
	INNER JOIN books ON user_books.book_id=books.id`)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return repo.getUserBooks(res)
}

func (repo *sqlUserBookRepository) FindAllByUsername(username string) ([]entity.UserBook, error) {
	res, err := repo.db.Query(`SELECT books.id, books.title, books.author, user_books.username 
	FROM user_books 
	INNER JOIN books ON user_books.book_id=books.id 
	WHERE user_books.username=?`, username)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return repo.getUserBooks(res)
}

func (repo *sqlUserBookRepository) Find(username, bookId string) (*entity.UserBook, error) {
	var book entity.UserBook
	row := repo.db.QueryRow(`SELECT books.id, books.title, books.author, user_books.username 
	FROM user_books 
	INNER JOIN books ON user_books.book_id=books.id 
	WHERE user_books.book_id=? AND user_books.username=?`, bookId, username)
	err := row.Scan(&book.BookId, &book.Title, &book.Author, &book.Username)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (repo *sqlUserBookRepository) Delete(username, bookId string) error {
	res, err := repo.db.Exec("DELETE FROM user_books WHERE username=? AND book_id=?", username, bookId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return custerror.UserBookNotFoundError(username, bookId)
	}
	return nil
}

func (repo *sqlUserBookRepository) Save(username, bookId string) error {
	_, err := repo.db.Exec("INSERT INTO user_books (book_id,username) VALUES(?, ?)", bookId, username)
	if err != nil {
		return err
	}
	return nil
}

func (repo *sqlUserBookRepository) getUserBooks(rows *sql.Rows) ([]entity.UserBook, error) {
	result := make([]entity.UserBook, 0)
	for rows.Next() {
		var book entity.UserBook
		if err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.Username); err != nil {
			return nil, err
		}
		result = append(result, book)
	}
	return result, nil
}
