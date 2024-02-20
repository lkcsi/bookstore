package entity

type UserBook struct {
	Username string `json:"username" binding:"required"`
	BookId   string `json:"book-id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
}
