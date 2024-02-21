CREATE DATABASE book_db;
use book_db;

CREATE TABLE books (
	id VARCHAR(255) PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	author VARCHAR(255) NOT NULL,
	quantity INT NOT NULL
);

INSERT INTO books (id, title, author, quantity)
VALUES 
	('bdecc1de-8b7d-4bf0-8154-f29d22b72be4', 'Title_1', 'Author_1', 0),
	('5d7d1e49-4183-4489-8646-8711c113b672', 'Title_2', 'Author_2', 1),
	('a4d5396b-dd25-499e-93f7-836a41772ba6', 'Title_3', 'Author_3', 2),
	('694606c3-671c-4297-9a8b-1b87f39b8422', 'Title_4', 'Author_4', 3);

CREATE TABLE users (
	username VARCHAR(255) PRIMARY KEY,
	password VARCHAR(255) NOT NULL
);

INSERT INTO users (username, password)
VALUES 
	('test', '$2a$10$KJKTiTcOhHjVIVH74u8pCOv18tzOs4Fd8bd8Dl7mZlJy/q2Tj2Vjq');


CREATE TABLE user_books (
	username VARCHAR(255) NOT NULL,
	book_id VARCHAR(255) NOT NULL,
	PRIMARY KEY(username, book_id),
	FOREIGN KEY(username) REFERENCES users(username),
	FOREIGN KEY(book_id) REFERENCES books(id)
);
