package rosters

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	{
		apiV1 := router.Group("/api/v1/rosters")

		apiV1.GET("/", getAllRosters)
	}
}
