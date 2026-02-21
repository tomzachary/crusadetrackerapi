package factions

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	{
		apiV1 := router.Group("/api/v1/factions")

		apiV1.GET("/", getAllFactions)
	}
}
