package factions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllFactions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}
