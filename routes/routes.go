package routes

import (
	"go-cron-job-microservice/controllers"

	"github.com/gin-gonic/gin"
)

// Initialize registers all REST endpoints to the router.
func Initialize(router *gin.Engine) {
	router.POST("/add", controllers.CronJob.Add)

	router.GET("/status/:jobid", controllers.CronJob.GetStatus)

	router.DELETE("/delete/:id", controllers.CronJob.Delete)
}
