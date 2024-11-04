package moviehdl

import (
	"net/http"

	"github.com/EliabeBastosDias/cinema-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func (mh *MovieHandler) ListMoviesAction(context *gin.Context) {
	var requestBody getMovieDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	movies, err := mh.service.List()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": movies})
}
