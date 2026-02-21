package armies

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	{
		apiV1 := router.Group("/api/v1/armies")

		apiV1.GET("/", getAllArmies)
	}
}
