package cronjobmanager

import (
	"go-cron-job-microservice/models"
	cronjob "go-cron-job-microservice/services/cronJob"
	sendrequest "go-cron-job-microservice/services/sendRequest"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var cronJobScheduler = cron.New()

// CronJobManager contains all functions to manage cron jobs.
type CronJobManager models.CronJobManager

var cm = CronJobManager{
	CachedJobs: make(map[primitive.ObjectID]models.CronJob),
}

// Start loads all cron jobs from the database and queues them.
func (cm *CronJobManager) Start() {
	allJobsFromDB, err := cm.loadJobsFromDB()
	if err != nil {
		fmt.Println("Aborting cron job launch")
		return
	}

	cronJobScheduler.Start()

	for _, cronJob := range allJobsFromDB {
		cm.launchCronJob(cronJob)
	}

	log.Printf("✅ Successfully started %v cron jobs\n", len(allJobsFromDB))
}

// Add persists a cron job in the database and queues it.
func (cm *CronJobManager) Add(cronJob models.CronJob) (primitive.ObjectID, error) {
	databaseID, err := cronjob.SaveToDatabase(cronJob)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	cronJob.ID = databaseID

	cm.launchCronJob(&cronJob)

	return databaseID, nil
}

// Remove removes a cron job from the queue and the database.
func (cm *CronJobManager) Remove(jobID primitive.ObjectID) models.CronJobDeletedResponse {
	var cronJobStatus models.CronJobDeletedResponse

	cronJobStatus.Ok = true
	cronJobStatus.Stopped = true
	cronJobStatus.DeletedFromDatabase = true

	if cronJob, ok := cm.CachedJobs[jobID]; ok {
		cronJobScheduler.Remove(cronJob.CronJobShedulerID)
		delete(cm.CachedJobs, jobID)
	} else {
		cronJobStatus.Stopped = false
	}

	if err := cronjob.DeleteFromDB(bson.M{"_id": jobID}); err != nil {
		cronJobStatus.Ok = false
		cronJobStatus.DeletedFromDatabase = false
		cronJobStatus.ErrorMessage = err.Error()
	}

	return cronJobStatus
}

// GetStatus returns the status for the cron job with the given ID.
func (cm *CronJobManager) GetStatus(jobID primitive.ObjectID) models.CronJobStatusResponse {
	var cronJobStatus models.CronJobStatusResponse

	cronJobStatus.Ok = true
	cronJobStatus.LoadedAndRunning = false

	if _, ok := cm.CachedJobs[jobID]; ok {
		cronJobStatus.LoadedAndRunning = true
	}

	if exists, err := cronjob.CheckIfExists(jobID); err == nil {
		cronJobStatus.PersistedInDatabase = exists
	} else {
		cronJobStatus.Ok = false
		cronJobStatus.PersistedInDatabase = false
		cronJobStatus.ErrorMessage = err.Error()
	}

	return cronJobStatus
}

func (cm *CronJobManager) loadJobsFromDB() ([]*models.CronJob, error) {
	allCronJobs, err := cronjob.LoadAllFromDB()

	if err != nil {
		return []*models.CronJob{}, err
	}

	return allCronJobs, nil
}

func (cm *CronJobManager) launchCronJob(cronJob *models.CronJob) {
	cronID, err := cronJobScheduler.AddFunc(cronJob.Frequency, func() {
		url := cronJob.URL

		if cronJob.Parameter != "" {
			url += cronJob.Parameter
		}

		res, err := sendrequest.Post(cronJob.Headers, url, cronJob.Payload)
		if err != nil || res.StatusCode != 200 {
			return
		}

		timestamp := time.Now()
		cronjob.FindOneAndUpdate(bson.M{"_id": cronJob.ID}, bson.M{"$set": bson.M{"lastExecution": timestamp}})
	})
	if err != nil {
		return
	}

	cronJob.CronJobShedulerID = cronID

	cm.CachedJobs[cronJob.ID] = *cronJob
}

// GetManager returns the CronJobManager instance.
func GetManager() *CronJobManager {
	return &cm
}
