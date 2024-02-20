package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lkcsi/bookstore/custerror"
)

func setApiError(context *gin.Context, err error) {
	switch e := err.(type) {
	case validator.ValidationErrors:
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": getValidationMsg(e[0])})
	case custerror.CustError:
		s := err.(custerror.CustError)
		context.AbortWithStatusJSON(s.Code(), gin.H{"error": s.Error()})
	default:
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func getValidationMsg(fe validator.FieldError) string {
	switch fe.ActualTag() {
	case "required":
		return fe.Field() + " field is mandatory"
	case "gte":
		return fe.Field() + " must be greater than or equals " + fe.Param()
	}
	return "unkown error"
}
