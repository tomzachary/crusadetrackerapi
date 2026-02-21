package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	{
		apiV1 := router.Group("/api/v1/users")

		apiV1.GET("/", getAllUsers)
	}
}
