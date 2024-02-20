package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lkcsi/bookstore/controller"
	"github.com/lkcsi/bookstore/repository"
	"github.com/lkcsi/bookstore/service"
)

var bookRepository repository.BookRepository
var userRepository repository.UserRepository
var userBookRepository repository.UserBookRepository
var userService service.UserService
var bookService service.BookService
var userBookService service.UserBookService

func authEnabled() bool {
	v := os.Getenv("AUTH_ENABLED")
	r, ok := strconv.ParseBool(v)
	if ok == nil {
		return r
	}
	return false
}

func initServices() {
	if os.Getenv("BOOKS_REPOSITORY") == "SQL" {

		bookRepository = repository.SqlBookRepository()
		userRepository = repository.SqlUserRepository()
	} else {
		database := repository.NewImDatabase()
		bookRepository = repository.ImBookRepository(database)
		userBookRepository = repository.ImUserBookRepository(database)
		userRepository = repository.ImUserRepository()
	}

	bookService = service.NewBookService(&bookRepository)
	userService = service.NewUserService(&userRepository)
	userBookService = service.NewUserBookService(&userBookRepository, &bookRepository)
}

func getAuthService() service.AuthService {
	if authEnabled() {
		return service.JwtAuthService()
	}
	return service.FakeAuthService()
}

func main() {
	godotenv.Load()
	initServices()

	server := gin.Default()

	bookApiController := controller.NewBookApiController(&bookService)
	userController := controller.NewUserController(&userService)
	userBookController := controller.NewUserBookController(&userBookService)

	authService := getAuthService()

	books := server.Group("/api/books")
	books.Use(authService.HeaderAuth)
	books.GET("", bookApiController.FindAll)
	books.GET("/:id", bookApiController.FindById)
	books.DELETE("/:id", bookApiController.DeleteBookById)
	books.DELETE("", bookApiController.DeleteAll)
	books.POST("", bookApiController.Save)

	userBooks := server.Group("/api/user-books")
	userBooks.Use(authService.HeaderAuth)
	userBooks.PATCH("/:username/checkout/:id", userBookController.Checkout)
	userBooks.PATCH("/:username/return/:id", userBookController.Return)
	userBooks.GET("", userBookController.FindAll)
	userBooks.GET("/:username", userBookController.FindAllByUsername)

	users := server.Group("/api/users")
	users.GET("/:username", userController.FindByUsername)
	users.POST("", userController.Save)
	users.POST("/login", userController.Login)

	server.GET("/api/health-check", healthCheck)

	server.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("BOOKS_API_PORT")))
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "OK"})
}
