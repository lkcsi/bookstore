package entity

type Book struct {
	Id       string `json:"id" binding:"-"`
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Quantity *int   `json:"quantity" binding:"required,gte=0"`
}
