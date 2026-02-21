package factions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	{
		apiV1 := router.Group("/api/v1/factions")

		apiV1.GET("/", getAllFactions)
	}
}
func getAllFactions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}
