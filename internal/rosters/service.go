package rosters

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllRosters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}
