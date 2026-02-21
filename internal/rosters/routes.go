package rosters

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	{
		apiV1 := router.Group("/api/v1/rosters")

		apiV1.GET("/", getAllRosters)
	}
}
func getAllRosters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}
