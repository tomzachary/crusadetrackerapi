package armies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllArmies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}
