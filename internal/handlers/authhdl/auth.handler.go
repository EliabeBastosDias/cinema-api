package authhdl

import (
	"github.com/EliabeBastosDias/cinema-api/internal/adapters"
	authservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/auth"
	jwthandler "github.com/EliabeBastosDias/cinema-api/internal/core/utils/jwt"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service authservice.AuthService
	logger  logger.Provider
}

func (cc *AuthHandler) SetUpRoutes(r *gin.Engine) {
	r.POST("/login", cc.LoginAction)
	r.POST("/users/register", cc.RegisterUserAction)
}

func NewHandler(apt *adapters.Adapters) *AuthHandler {
	return &AuthHandler{
		service: *authservice.NewAuthService(apt.Repositories, jwthandler.JWTHandler{}, apt.Logger),
		logger:  apt.Logger,
	}
}
