package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lkcsi/bookstore/custerror"
)

func SetError(context *gin.Context, err error) {
	switch err.(type) {
	case validator.ValidationErrors:
		s := err.(validator.ValidationErrors)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": getErrorMsg(s[0])})
	case custerror.CustError:
		s := err.(custerror.CustError)
		context.AbortWithStatusJSON(s.Code(), gin.H{"error": s.Error()})
	default:
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.ActualTag() {
	case "required":
		return fe.Field() + " field is mandatory"
	case "gte":
		return fe.Field() + " must be greater than or equals " + fe.Param()
	}
	return "unkown error"
}
