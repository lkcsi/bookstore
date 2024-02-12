package entity

type Book struct {
	Id       string `json:"id" binding:"-"`
	Title    string `json:"title" form:"title" binding:"required"`
	Author   string `json:"author" form:"author" binding:"required"`
	Quantity *int   `json:"quantity" form:"quantity" binding:"required,gte=0"`
}
