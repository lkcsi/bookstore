package controller

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/lkcsi/bookstore/entity"
	"github.com/lkcsi/bookstore/service"
)

type MyBookView interface {
	Return(context *gin.Context)
	Get(context *gin.Context)
}

type myBookView struct {
	bookService service.BookService
}

func NewMyBookView(s *service.BookService) *myBookView {
	return &myBookView{bookService: *s}
}

func (b *myBookView) Get(context *gin.Context) {
	tmpl, _ := template.New("").ParseFiles("template/index.html", "template/my-books.html")
	books, err := b.bookService.FindAll()
	if err != nil {
		books = make([]entity.Book, 0)
	}

	tmpl.ExecuteTemplate(context.Writer, "index", gin.H{
		"Books": books,
	})
}

func (b *myBookView) Return(context *gin.Context) {
	id := context.Param("id")
	book, err := b.bookService.Return(id)
	if err != nil {
		setViewError(context, err)
	}

	tmpl, _ := template.New("t").Parse(itemHtml(book))
	tmpl.Execute(context.Writer, nil)
}
