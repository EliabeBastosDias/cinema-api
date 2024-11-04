package moviehdl

import (
	"github.com/EliabeBastosDias/cinema-api/internal/adapters"
	movieservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/movie"
	jwthandler "github.com/EliabeBastosDias/cinema-api/internal/core/utils/jwt"
	"github.com/EliabeBastosDias/cinema-api/internal/middlewares"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service    movieservice.MovieService
	logger     logger.Provider
	jwtHandler jwthandler.JWTHandler
}

func (cc *MovieHandler) SetUpRoutes(r *gin.Engine) {
	movieRoutes := r.Group("/movies")

	movieRoutes.POST("/", middlewares.AuthMiddleware(cc.jwtHandler, cc.logger), cc.CreateMovieAction)
	movieRoutes.GET("/:movieToken", cc.GetMovieAction)
	movieRoutes.GET("/", cc.ListMoviesAction)
	movieRoutes.PUT("/:movieToken", middlewares.AuthMiddleware(cc.jwtHandler, cc.logger), cc.UpdateMovieAction)
	movieRoutes.POST("/:movieToken/inactivate", middlewares.AuthMiddleware(cc.jwtHandler, cc.logger), cc.UpdateMovieAction)
}

func NewHandler(apt *adapters.Adapters) *MovieHandler {
	return &MovieHandler{
		service:    *movieservice.New(apt.Repositories, apt.Logger),
		logger:     apt.Logger,
		jwtHandler: jwthandler.JWTHandler{},
	}
}
