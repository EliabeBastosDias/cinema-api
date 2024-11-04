package moviehdl

import (
	"net/http"

	movieservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/movie"
	"github.com/EliabeBastosDias/cinema-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type createMovieDto struct {
	Name     string `json:"name" binding:"required"`
	Director string `json:"director" binding:"required"`
	Duration int    `json:"duration" binding:"required"`
}

func (mh *MovieHandler) CreateMovieAction(context *gin.Context) {
	var requestBody createMovieDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	params := movieservice.CreateMovieParams{
		Name:     requestBody.Name,
		Director: requestBody.Director,
		Duration: requestBody.Duration,
	}

	createdMovie, err := mh.service.Create(params)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": createdMovie})
}
