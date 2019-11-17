package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendJSON returns a response in JSON to the client.
func SendJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
