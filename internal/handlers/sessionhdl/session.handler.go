package sessionhdl

import (
	"github.com/EliabeBastosDias/cinema-api/internal/adapters"
	sessionservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/session"
	jwthandler "github.com/EliabeBastosDias/cinema-api/internal/core/utils/jwt"
	"github.com/EliabeBastosDias/cinema-api/internal/middlewares"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	service    sessionservice.SessionService
	logger     logger.Provider
	jwtHandler jwthandler.JWTHandler
}

func (cc *SessionHandler) SetUpRoutes(r *gin.Engine) {
	sessionRoutes := r.Group("/sessions")

	sessionRoutes.POST("/", middlewares.AuthMiddleware(cc.jwtHandler, cc.logger), cc.CreateSessionAction)
	sessionRoutes.GET("/:sessionToken", cc.GetSessionAction)
	sessionRoutes.GET("/", cc.ListSessionsAction)
	sessionRoutes.PUT("/:sessionToken", middlewares.AuthMiddleware(cc.jwtHandler, cc.logger), cc.UpdateSessionAction)
	sessionRoutes.POST("/:sessionToken/inactivate", middlewares.AuthMiddleware(cc.jwtHandler, cc.logger), cc.UpdateSessionAction)
}

func NewHandler(apt *adapters.Adapters) *SessionHandler {
	return &SessionHandler{
		service:    *sessionservice.New(apt.Repositories, apt.Logger),
		logger:     apt.Logger,
		jwtHandler: jwthandler.JWTHandler{},
	}
}
