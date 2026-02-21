package armies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) RegisterRoutes(router *gin.Engine) {
	service := Service{}
	{
		apiV1 := router.Group("/api/v1/armies")

		apiV1.GET("/", service.handleGetAllArmies)
	}
}

func (s *Service) handleGetAllArmies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, s.GetAllArmies())
}
