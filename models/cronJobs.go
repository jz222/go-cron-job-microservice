package models

import (
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CronJobManager manages the cron jobs.
type CronJobManager struct {
	CachedJobs map[primitive.ObjectID]CronJob
}

// CronJob contains the information about the cron job.
type CronJob struct {
	ID                primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Frequency         string                 `json:"frequency" bson:"frequency"`
	URL               string                 `json:"url" bson:"url"`
	Headers           map[string]string      `json:"headers" bson:"headers"`
	Parameter         string                 `json:"parameter" bson:"parameter"`
	Payload           map[string]interface{} `json:"payload" bson:"payload"`
	LastExecution     time.Time              `json:"lastExecution" bson:"lastExecution"`
	CronJobShedulerID cron.EntryID           `json:"cronJobShedulerId" bson:"cronJobShedulerId"`
}

// CronJobStatusResponse contains the status of a cron job.
type CronJobStatusResponse struct {
	Ok                  bool   `json:"ok"`
	LoadedAndRunning    bool   `json:"loadedAndRunning"`
	PersistedInDatabase bool   `json:"persistedInDatabase"`
	ErrorMessage        string `json:"errorMessage,omitempty"`
}

// CronJobDeletedResponse contains the status of a deleted cron job.
type CronJobDeletedResponse struct {
	Ok                  bool   `json:"ok"`
	Stopped             bool   `json:"stopped"`
	DeletedFromDatabase bool   `json:"deletedFromDatabase"`
	ErrorMessage        string `json:"errorMessage,omitempty"`
}
