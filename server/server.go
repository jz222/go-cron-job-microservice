package server

import (
	"go-cron-job-microservice/keys"
	"go-cron-job-microservice/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	port   = fmt.Sprintf(":%s", keys.GetKeys().PORT)
	router *gin.Engine
)

// Start runs the server.
func Start() {
	router = gin.Default()

	routes.Initialize(router)

	if err := router.Run(port); err != nil {
		panic(err)
	}
}
