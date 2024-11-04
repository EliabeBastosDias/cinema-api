package healthy

import (
	"net/http"

	"github.com/EliabeBastosDias/cinema-api/internal/adapters"
	"github.com/gin-gonic/gin"
)

type Provider struct {
	a *adapters.Adapters
}

func NewHandler(a *adapters.Adapters) Provider {
	return Provider{a: a}
}

func (p Provider) SetUpRoutes(c *gin.Engine) {
	c.GET("/health", p.Health)
}

func (p Provider) Health(c *gin.Context) {
	if err := p.a.DB.Ping(); err != nil {
		p.a.Logger.Error("failed to ping db: err: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "database system unavailable",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "system running",
	})
}
