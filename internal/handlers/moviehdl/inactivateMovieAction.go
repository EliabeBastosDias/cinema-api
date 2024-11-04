package moviehdl

import (
	"net/http"

	movieservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/movie"
	"github.com/gin-gonic/gin"
)

func (cc *MovieHandler) InactivateMovieAction(context *gin.Context) {
	movieToken := context.Param("movieToken")

	params := movieservice.InactivateMovieParams{
		MovieToken: movieToken,
	}

	err := cc.service.Inactivate(params)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"success": true})
}
