package controllers

import (
	"go-cron-job-microservice/models"
	cronjobmanager "go-cron-job-microservice/services/cronJobManager"
	"go-cron-job-microservice/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cronJobControllers struct{}

// CronJob contains of the controllers related to cron jobs.
var CronJob cronJobControllers

// Add validates the request to ensure the given cron job contains a frequency and URL
// before persisting the cron job and adding it to the queue.
func (cj *cronJobControllers) Add(c *gin.Context) {
	var cronJob models.CronJob
	var errorResponse models.Error

	if err := json.NewDecoder(c.Request.Body).Decode(&cronJob); err != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "There was an error while parsing the request"
		fmt.Println(err)

		utils.SendError(c, errorResponse)
		return
	}

	if cronJob.Frequency == "" {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "Cron job frequency was not provided"

		utils.SendError(c, errorResponse)
		return
	}

	if cronJob.URL == "" {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "URL was not provided"

		utils.SendError(c, errorResponse)
		return
	}

	cronJob.LastExecution = time.Now()

	cronJobManager := cronjobmanager.GetManager()

	id, err := cronJobManager.Add(cronJob)
	if err != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = err.Error()

		utils.SendError(c, errorResponse)
		return
	}

	utils.SendJSON(c, map[string]string{"cronID": id.Hex()})
}

// Delete validates the request to ensure a valid cron job ID is present and removes
// a cron job with the given ID from the queue and the database.
func (cj *cronJobControllers) Delete(c *gin.Context) {
	var errorResponse models.Error

	cronJobIDParameter := c.Param("id")

	cronJobID, err := primitive.ObjectIDFromHex(cronJobIDParameter)
	if err != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "The provided cron job ID is invalid"

		utils.SendError(c, errorResponse)
		return
	}

	cronJobManager := cronjobmanager.GetManager()
	status := cronJobManager.Remove(cronJobID)

	utils.SendJSON(c, status)
}

// GetStatus validates the request to ensure a valid cron job ID is
// present and checks the status of the cron job with the given ID.
func (cj *cronJobControllers) GetStatus(c *gin.Context) {
	var errorResponse models.Error

	cronJobIDParamter := c.Param("jobid")

	cronJobID, err := primitive.ObjectIDFromHex(cronJobIDParamter)
	if err != nil {
		errorResponse.Code = http.StatusBadRequest
		errorResponse.Message = "The provided cron job ID is invalid"

		utils.SendError(c, errorResponse)
		return
	}

	cronJobManager := cronjobmanager.GetManager()
	cronJobStatus := cronJobManager.GetStatus(cronJobID)

	utils.SendJSON(c, cronJobStatus)
}
