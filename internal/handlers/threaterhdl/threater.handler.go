package threaterhdl

import (
	"github.com/EliabeBastosDias/cinema-api/internal/adapters"
	threaterservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/threater"
	jwthandler "github.com/EliabeBastosDias/cinema-api/internal/core/utils/jwt"
	"github.com/EliabeBastosDias/cinema-api/internal/middlewares"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type ThreaterHandler struct {
	service    threaterservice.ThreaterService
	logger     logger.Provider
	jwtHandler jwthandler.JWTHandler
}

func (cc *ThreaterHandler) SetUpRoutes(r *gin.Engine) {
	threaterRoutes := r.Group("/threaters")

	threaterRoutes.POST("/", middlewares.AuthMiddleware(cc.jwtHandler, cc.logger), cc.CreateThreaterAction)
	threaterRoutes.GET("/:threaterToken", cc.GetThreaterAction)
	threaterRoutes.GET("/", cc.ListThreatersAction)
	threaterRoutes.PUT("/:threaterToken", middlewares.AuthMiddleware(cc.jwtHandler, cc.logger), cc.UpdateThreaterAction)
	threaterRoutes.POST("/:threaterToken/inactivate", middlewares.AuthMiddleware(cc.jwtHandler, cc.logger), cc.UpdateThreaterAction)
}

func NewHandler(apt *adapters.Adapters) *ThreaterHandler {
	return &ThreaterHandler{
		service:    *threaterservice.New(apt.Repositories, apt.Logger),
		logger:     apt.Logger,
		jwtHandler: jwthandler.JWTHandler{},
	}
}
