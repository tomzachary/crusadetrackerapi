package armies

import (
	"crusadetrackerapi/internal/common"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Service) RegisterRoutes(router *gin.Engine) {
	apiV1 := router.Group("/api/v1/armies")

	apiV1.GET("/", s.handleGetAllArmies)
	apiV1.POST("/", s.handlePostArmy)
	apiV1.PUT("/:id", s.handlePutArmy)
	apiV1.DELETE("/:id", s.handleDeleteArmy)
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

	if army, err := s.CreateArmy(newArmy); err != nil {
		errorString := fmt.Sprintf("Error creating army: %v\n", err)
		fmt.Print(errorString)
		c.String(http.StatusInternalServerError, errorString)
		return
	} else {
		c.IndentedJSON(http.StatusOK, army)
	}
}

func (s *Service) handlePutArmy(c *gin.Context) {
	idStr, ok := c.Params.Get("id")
	if !ok {
		c.String(http.StatusBadRequest, "Error: no id given to PUT request")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error parsing id: %v\n", err))
		return
	}

	var updatedArmy Army
	if err := common.ParseBody(c, &updatedArmy); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error parsing data: %v\n", err))
		return
	}

	if army, err := s.UpdateArmy(id, updatedArmy); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error updating army: %v\n", err))
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
