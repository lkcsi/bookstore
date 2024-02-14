package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lkcsi/bookstore/custerror"
)

func setApiError(context *gin.Context, err error) {
	switch err.(type) {
	case validator.ValidationErrors:
		s := err.(validator.ValidationErrors)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": getValidationMsg(s[0])})
	case custerror.CustError:
		s := err.(custerror.CustError)
		context.AbortWithStatusJSON(s.Code(), gin.H{"error": s.Error()})
	default:
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func setViewError(context *gin.Context, err error) {
	context.Writer.Header().Add("HX-Retarget", "#errors")
	errors := ""

	switch err.(type) {
	case validator.ValidationErrors:
		s := err.(validator.ValidationErrors)
		for _, elem := range s {
			errors = fmt.Sprintf("%s %s", errors, getDangerMessage(getValidationMsg(elem)))
		}
	case custerror.CustError:
		errors = err.Error()
	default:
		errors = "Internal sever error"
	}

	tmpl, _ := template.New("t").Parse(errors)
	tmpl.Execute(context.Writer, nil)
}

func getDangerMessage(msg string) string {
	return fmt.Sprintf("<p class='alert alert-danger'>%s</p>", msg)
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
