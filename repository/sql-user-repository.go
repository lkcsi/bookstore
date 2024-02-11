package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lkcsi/bookstore/entity"
)

type sqlUserRepository struct {
	db *sql.DB
}

func SqlUserRepository() UserRepository {
	pwd := os.Getenv("BOOKS_DB_PASSWORD")
	port := os.Getenv("BOOKS_DB_PORT")
	host := os.Getenv("BOOKS_DB_HOST")
	conn := fmt.Sprintf("root:%s@tcp(%s:%s)/book_db", pwd, host, port)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println("Unable to connect")
	}
	return &sqlUserRepository{db}
}

func (repo *sqlUserRepository) Save(userRequest *entity.User) error {
	_, err := repo.db.Exec("INSERT INTO users (username, password) VALUES(?, ?)", userRequest.Username, userRequest.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *sqlUserRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	row := repo.db.QueryRow("SELECT username, password FROM users WHERE username=?", username)
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
