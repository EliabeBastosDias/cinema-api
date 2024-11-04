package threaterhdl

import (
	"net/http"

	"github.com/EliabeBastosDias/cinema-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func (mh *ThreaterHandler) ListThreatersAction(context *gin.Context) {
	var requestBody getThreaterDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	threaters, err := mh.service.List()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": threaters})
}
