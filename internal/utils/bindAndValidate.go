package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type errorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Size should be at least " + fe.Param()
	case "datetime":
		return fmt.Sprintf("field %v: invalid field format: expected format %v", fe.Field(), fe.Param())
	case "email":
		return "Invalid email format"
	}
	return "Unknown error"
}

func BindAndValidate(context *gin.Context, requestBody interface{}) error {
	if err := context.BindJSON(requestBody); err != nil {
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			out := make([]errorMsg, len(ve))

			for i, fe := range ve {
				out[i] = errorMsg{fe.Field(), getErrorMsg(fe)}
			}
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request body", "errors": out})
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		}
		return err
	}
	return nil
}
