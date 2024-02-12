package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lkcsi/bookstore/controller"
	"github.com/lkcsi/bookstore/entity"
	"github.com/lkcsi/bookstore/service"
)

func authEnabled() bool {
	v := os.Getenv("AUTH_ENABLED")
	r, ok := strconv.ParseBool(v)
	if ok == nil {
		return r
	}
	return false
}

var authService service.AuthService
var bookService service.BookService

func getAuthService() service.AuthService {
	if authEnabled() {
		return service.JwtAuthService()
	}
	return service.FakeAuthService()
}

func getBookService() service.BookService {
	if os.Getenv("BOOKS_REPOSITORY") == "SQL" {
		return service.NewSqlBookService()
	}
	return service.InMemoryBookService()
}

func getUserService() service.UserService {
	return service.SqlUserService()
}

func main() {
	godotenv.Load()

	server := gin.Default()

	bookService = getBookService()
	bookApiController := controller.NewBookApiController(&bookService)
	bookViewController := controller.NewBookViewController(&bookService)
	authService = getAuthService()

	server.GET("/index", mainPage)
	server.POST("/add-book", bookViewController.Save)

	books := server.Group("/books")
	books.Use(authService.Auth)
	books.GET("", bookApiController.FindAll)
	books.GET("/:id", bookApiController.FindById)
	books.DELETE("/:id", bookApiController.DeleteBookById)
	books.DELETE("", bookApiController.DeleteAll)
	books.POST("", bookApiController.Save)
	books.PATCH("/:id/checkout", bookApiController.CheckoutBook)

	userService := getUserService()
	userController := controller.NewUserController(&userService)

	users := server.Group("/users")
	users.GET("/:username", userController.FindByUsername)
	users.POST("", userController.Save)
	users.POST("/login", userController.Login)

	health := server.Group("/health-check")
	health.GET("", healthCheck)

	server.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("BOOKS_API_PORT")))
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "O.K."})
}

func mainPage(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	books, err := bookService.FindAll()
	if err != nil {
		books = make([]entity.Book, 0)
	}

	tmpl.Execute(c.Writer, gin.H{
		"title": "Main Title",
		"Books": books,
	})
}
