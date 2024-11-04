package moviehdl

import (
	"net/http"

	movieservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/movie"
	"github.com/EliabeBastosDias/cinema-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type updateMovieDto struct {
	Name     string `json:"name" binding:"required"`
	Director string `json:"director" binding:"required"`
	Duration int    `json:"duration" binding:"required"`
}

func (cc *MovieHandler) UpdateMovieAction(context *gin.Context) {
	movieToken := context.Param("movieToken")
	var requestBody updateMovieDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	params := movieservice.UpdateMovieParams{
		MovieToken: movieToken,
		Name:       requestBody.Name,
		Director:   requestBody.Director,
		Duration:   requestBody.Duration,
	}

	err := cc.service.Update(params)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"success": true})
}
