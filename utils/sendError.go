package utils

import (
	"go-cron-job-microservice/models"

	"github.com/gin-gonic/gin"
)

// SendError returns an error in JSON to the client.
func SendError(c *gin.Context, err models.Error) {
	c.JSON(err.Code, err)
}
