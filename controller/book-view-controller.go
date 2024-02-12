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
}

type bookViewController struct {
	bookService service.BookService
}

func NewBookViewController(s *service.BookService) *bookViewController {
	return &bookViewController{bookService: *s}
}

func (c *bookViewController) Save(context *gin.Context) {
	var requestedBook entity.Book
	if err := context.ShouldBind(&requestedBook); err != nil {
		context.Writer.Header().Add("HX-Retarget", "#errors")
		context.Writer.Header().Add("HX-Reswap", "innerHTML")
		tmpl, _ := template.New("t").Parse(fmt.Sprintf("<p class='alert alert-danger'>%s</p>", err.Error()))
		tmpl.Execute(context.Writer, nil)
		context.AbortWithStatus(400)
		return
	}

	newBook, err := c.bookService.Save(requestedBook)
	if err != nil {
		return
	}
	htmlStr := fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%d</td></tr>",
		newBook.Title, newBook.Author, *newBook.Quantity)

	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(context.Writer, nil)

	//context.IndentedJSON(http.StatusCreated, newBook)
}
