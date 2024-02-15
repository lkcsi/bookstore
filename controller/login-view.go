package controller

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/lkcsi/bookstore/entity"
	"github.com/lkcsi/bookstore/service"
)

type LoginView interface {
	Get(context *gin.Context)
	Login(context *gin.Context)
}

type loginView struct {
	userService service.UserService
}

func NewLoginView(s *service.UserService) *loginView {
	return &loginView{userService: *s}
}

func (l *loginView) Get(context *gin.Context) {
	tmpl := template.Must(template.ParseFiles("template/login.html"))
	tmpl.Execute(context.Writer, nil)
}

func (l *loginView) Login(context *gin.Context) {
	var requestUser entity.User
	if err := context.ShouldBind(&requestUser); err != nil {
		setViewError(context, err)
		return
	}
	token, err := l.userService.Login(&requestUser)
	if err != nil {
		fmt.Println(err.Error())
		setViewError(context, err)
		return
	}
	fmt.Println("3")
	context.SetCookie("auth", token, (15 * 3600), "/", "localhost", false, true)
}
