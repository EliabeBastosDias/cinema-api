package threaterhdl

import (
	"net/http"

	threaterservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/threater"
	"github.com/gin-gonic/gin"
)

type getThreaterDto struct {
	ThreaterToken string `json:"threater_token" binding:"required"`
}

func (mh *ThreaterHandler) GetThreaterAction(context *gin.Context) {
	threaterToken := context.Param("threaterToken")

	threater, err := mh.service.Get(threaterservice.GetThreaterParams{
		ThreaterToken: threaterToken,
	})
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": threater})
}
