package controller

import (
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
		setViewError(context, err)
		return
	}
	context.SetCookie("auth", token, (1 * 60), "/", "localhost", false, true)
	context.Writer.Header().Add("HX-Redirect", "/index")
}
