package sessionhdl

import (
	"net/http"

	sessionservice "github.com/EliabeBastosDias/cinema-api/internal/core/use-cases/session"
	"github.com/gin-gonic/gin"
)

func (cc *SessionHandler) InactivateSessionAction(context *gin.Context) {
	sessionToken := context.Param("sessionToken")

	params := sessionservice.InactivateSessionParams{
		SessionToken: sessionToken,
	}

	err := cc.service.Inactivate(params)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"success": true})
}
