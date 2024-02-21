package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkcsi/bookstore/service"
)

type UserBookController interface {
	Return(context *gin.Context)
	FindAll(context *gin.Context)
	Find(context *gin.Context)
	FindAllByUsername(context *gin.Context)
	Checkout(context *gin.Context)
}

type userBookController struct {
	userBookService service.UserBookService
}

func NewUserBookController(s *service.UserBookService) UserBookController {
	return &userBookController{userBookService: *s}
}

func (c *userBookController) FindAll(context *gin.Context) {
	books, err := c.userBookService.FindAll()
	if err != nil {
		setApiError(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, books)
}

func (c *userBookController) Find(context *gin.Context) {
	username := context.Param("username")
	id := context.Param("id")
	book, err := c.userBookService.Find(username, id)
	if err != nil {
		setApiError(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}

func (c *userBookController) FindAllByUsername(context *gin.Context) {
	username := context.Param("username")
	userBook, err := c.userBookService.FindAllByUsername(username)
	if err != nil {
		setApiError(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, userBook)
}

func (c *userBookController) Checkout(context *gin.Context) {
	context.Writer.Header().Set("content-type", "application/json")

	username := context.Param("username")
	bookId := context.Param("id")

	resp, err := c.userBookService.Checkout(username, bookId)
	if err != nil {
		setApiError(context, err)
		return
	}

	context.IndentedJSON(http.StatusAccepted, resp)
}

func (c *userBookController) Return(context *gin.Context) {
	context.Writer.Header().Set("content-type", "application/json")

	username := context.Param("username")
	bookId := context.Param("id")

	err := c.userBookService.Return(username, bookId)
	if err != nil {
		setApiError(context, err)
		return
	}

	context.IndentedJSON(http.StatusAccepted, nil)
}
