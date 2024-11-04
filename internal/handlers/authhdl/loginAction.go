package authhdl

import (
	"net/http"

	authservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/auth"
	"github.com/gin-gonic/gin"
)

func (cc *AuthHandler) LoginAction(c *gin.Context) {
	var loginParams authservice.LoginParams
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		cc.logger.Error("Invalid login parameters", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login parameters"})
		return
	}

	loginResult, err := cc.service.Login(loginParams)
	if err != nil {
		cc.logger.Error("Login failed", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  loginResult.User,
		"token": loginResult.Token,
	})
}
