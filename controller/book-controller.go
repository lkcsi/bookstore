package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkcsi/bookstore/entity"
	"github.com/lkcsi/bookstore/service"
)

type BookApiController interface {
	FindAll(context *gin.Context)
	FindById(context *gin.Context)
	DeleteById(context *gin.Context)
	Save(context *gin.Context)
	DeleteAll(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookApiController(s *service.BookService) *bookController {
	return &bookController{bookService: *s}
}

func (c *bookController) FindAll(context *gin.Context) {
	books, err := c.bookService.FindAll()
	if err != nil {
		setApiError(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, books)
}

func (c *bookController) DeleteAll(context *gin.Context) {
	if err := c.bookService.DeleteAll(); err != nil {
		setApiError(context, err)
		return
	}
	context.IndentedJSON(http.StatusNoContent, "")
}

func (c *bookController) Save(context *gin.Context) {
	context.Writer.Header().Set("content-type", "application/json")
	var requestedBook entity.Book
	if err := context.BindJSON(&requestedBook); err != nil {
		setApiError(context, err)
		return
	}

	newBook, err := c.bookService.Save(requestedBook)
	if err != nil {
		setApiError(context, err)
		return
	}

	context.IndentedJSON(http.StatusCreated, newBook)
}

func (c *bookController) DeleteBookById(context *gin.Context) {
	id := context.Param("id")
	if err := c.bookService.DeleteById(id); err != nil {
		setApiError(context, err)
		return
	}
	context.IndentedJSON(http.StatusNoContent, nil)
}

func (c *bookController) FindById(context *gin.Context) {
	id := context.Param("id")
	book, err := c.bookService.FindById(id)
	if err != nil {
		setApiError(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}

type ErrorMsg struct {
	Message string `json:"message"`
}
