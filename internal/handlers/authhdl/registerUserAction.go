package authhdl

import (
	"net/http"

	authservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/auth"
	"github.com/gin-gonic/gin"
)

func (cc *AuthHandler) RegisterUserAction(c *gin.Context) {
	var params authservice.RegisterUserParams

	if err := c.ShouldBindJSON(&params); err != nil {
		cc.logger.Error("Invalid registration parameters", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration parameters"})
		return
	}

	user, err := cc.service.RegisterUser(params)
	if err != nil {
		cc.logger.Error("User registration failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
}
