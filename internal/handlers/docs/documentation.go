package docs

import (
	"net/http"

	"github.com/EliabeBastosDias/cinema-api/internal/adapters"
	"github.com/gin-gonic/gin"
)

var SWAGGUER_PATH = "internal/docs/swaguer.yaml"

type Provider struct {
	a           *adapters.Adapters
	swaggerPath string
}

func NewHandler(apt *adapters.Adapters) Provider {
	return Provider{a: apt, swaggerPath: SWAGGUER_PATH}
}

func (p Provider) SetUpRoutes(c *gin.Engine) {
	c.GET("/docs", p.ServeSwagger)
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

func (p Provider) ServeSwagger(c *gin.Context) {
	c.File(p.swaggerPath)
}
