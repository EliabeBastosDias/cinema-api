package threaterhdl

import (
	"net/http"

	threaterservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/threater"
	"github.com/gin-gonic/gin"
)

func (cc *ThreaterHandler) InactivateMovieAction(context *gin.Context) {
	threaterToken := context.Param("threaterToken")

	params := threaterservice.InactivateThreaterParams{
		ThreaterToken: threaterToken,
	}

	err := cc.service.Inactivate(params)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"success": true})
}
