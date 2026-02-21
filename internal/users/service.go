package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}
