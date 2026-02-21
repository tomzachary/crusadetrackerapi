package armies

import (
	"crusadetrackerapi/internal/common"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Service) RegisterRoutes(router *gin.Engine) {
	service := Service{}
	{
		apiV1 := router.Group("/api/v1/armies")

		apiV1.GET("/", service.handleGetAllArmies)
		apiV1.POST("/", service.handlePostArmy)
		apiV1.DELETE("/:id", service.handleDeleteArmy)
	}
}

func (s *Service) handleGetAllArmies(c *gin.Context) {
	if armies, err := s.GetAllArmies(); err != nil {
		errorString := fmt.Sprintf("Error getting All Armies: %v\n", err)
		fmt.Print(errorString)
		c.String(http.StatusInternalServerError, errorString)

	} else {
		c.IndentedJSON(http.StatusOK, armies)
	}

}

func (s *Service) handlePostArmy(c *gin.Context) {
	var newArmy Army
	if err := common.ParseBody(c, &newArmy); err != nil {
		errorString := fmt.Sprintf("Error parsing data: %v\n", err)
		fmt.Print(errorString)
		c.String(http.StatusInternalServerError, errorString)
		return
	}

	if army, err := s.CreateOrUpdateArmy(newArmy); err != nil {
		errorString := fmt.Sprintf("Error creating army: %v\n", err)
		fmt.Print(errorString)
		c.String(http.StatusInternalServerError, errorString)
		return
	} else {
		c.IndentedJSON(http.StatusOK, army)
	}
}

func (s *Service) handleDeleteArmy(c *gin.Context) {
	idToDelete, ok := c.Params.Get("id")
	if !ok {
		errorString := "Error no id given to Delete request"
		fmt.Print(errorString)
		c.String(http.StatusBadRequest, errorString)
		return
	}
	if id, err := strconv.Atoi(idToDelete); err != nil {
		errorString := fmt.Sprintf("Error Deleting Army: %v\n", err)
		fmt.Print(errorString)
		c.String(http.StatusBadRequest, errorString)
		return
	} else {
		s.DeleteArmy(id)
	}
}
