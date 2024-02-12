package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lkcsi/bookstore/entity"
	"github.com/lkcsi/bookstore/service"
)

type UserController interface {
	FindByUsername(*gin.Context)
	Save(*gin.Context)
	Login(*gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(s *service.UserService) UserController {
	return &userController{userService: *s}
}

func (u *userController) Save(c *gin.Context) {
	c.Writer.Header().Set("content-type", "application/json")
	var newUser entity.User
	if err := c.BindJSON(&newUser); err != nil {
		setApiError(c, err)
		return
	}
	if _, err := u.userService.Save(&newUser); err != nil {
		setApiError(c, err)
		return
	}
	c.IndentedJSON(201, newUser)
}

func (u *userController) FindByUsername(context *gin.Context) {
	username := context.Param("username")
	user, err := u.userService.FindByUsername(username)
	if err != nil {
		setApiError(context, err)
		return
	}
	context.IndentedJSON(200, user)
}

func (u *userController) Login(c *gin.Context) {
	c.Writer.Header().Set("content-type", "application/json")
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		setApiError(c, err)
		return
	}
	jwt, err := u.userService.Login(&user)
	if err != nil {
		setApiError(c, err)
		return
	}
	c.IndentedJSON(200, gin.H{"access_token": jwt})
}
