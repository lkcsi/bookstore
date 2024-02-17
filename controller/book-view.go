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
	Checkout(context *gin.Context)
	Get(context *gin.Context)
}

type bookView struct {
	bookService service.BookService
}

func NewBookView(s *service.BookService) *bookView {
	return &bookView{bookService: *s}
}

func checkoutButtonHtml(id string) string {
	return fmt.Sprintf("<button class='btn btn-success' hx-target='#book-%s' hx-post='/checkout-book/%s' hx-swap='outerHTML'> Checkout </button>", id, id)
}

func itemHtml(book *entity.Book) string {
	return fmt.Sprintf("<tr id='book-%s'><td>%s</td><td>%s</td><td>%d</td><td>%s</td></tr>",
		book.Id, book.Title, book.Author, *book.Quantity, checkoutButtonHtml(book.Id))
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

	tmpl, _ := template.New("t").Parse(itemHtml(newBook))
	tmpl.Execute(context.Writer, nil)
}

func (b *bookView) Get(context *gin.Context) {
	tmpl, _ := template.New("").ParseFiles("template/index.html", "template/books.html")
	books, err := b.bookService.FindAll()
	if err != nil {
		books = make([]entity.Book, 0)
	}

	tmpl.ExecuteTemplate(context.Writer, "index", gin.H{
		"Books": books,
	})
}

func (b *bookView) Checkout(context *gin.Context) {
	id := context.Param("id")
	book, err := b.bookService.Checkout(id)
	if err != nil {
		setViewError(context, err)
	}

	tmpl, _ := template.New("t").Parse(itemHtml(book))
	tmpl.Execute(context.Writer, nil)
}
