package common

import (
	"github.com/gin-gonic/gin"
)

func ParseBody[T any](c *gin.Context, newObject *T) error {
	if readBodyErr := c.ShouldBindJSON(&newObject); readBodyErr != nil {

		return readBodyErr
	}
	return nil
}
