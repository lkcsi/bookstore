package controller

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/lkcsi/bookstore/entity"
	"github.com/lkcsi/bookstore/service"
)

type BookView interface {
	Save(context *gin.Context)
	Get(context *gin.Context)
}

type bookView struct {
	bookService service.BookService
}

func NewBookView(s *service.BookService) *bookView {
	return &bookView{bookService: *s}
}

func (b *bookView) Save(context *gin.Context) {
	var requestedBook entity.Book
	if err := context.ShouldBind(&requestedBook); err != nil {
		setViewError(context, err)
		return
	}

	newBook, err := b.bookService.Save(requestedBook)
	if err != nil {
		return
	}
	htmlStr := fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%d</td></tr>",
		newBook.Title, newBook.Author, *newBook.Quantity)

	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(context.Writer, nil)
}

func (b *bookView) Get(context *gin.Context) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	books, err := b.bookService.FindAll()
	if err != nil {
		books = make([]entity.Book, 0)
	}

	tmpl.Execute(context.Writer, gin.H{
		"Books": books,
	})
}
