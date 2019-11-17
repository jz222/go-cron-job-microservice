package cronjob

import (
	"context"
	"go-cron-job-microservice/libs/mongodb"
	"go-cron-job-microservice/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// LoadAllFromDB loads all cron jobs from the database.
func LoadAllFromDB() ([]*models.CronJob, error) {
	var cronJobs []*models.CronJob

	collection := mongodb.GetClient().Collection("cronjob")
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Println("Failed to load all cron jobs from database")
		return cronJobs, err
	}

	for cur.Next(context.TODO()) {
		var cronJob models.CronJob

		err := cur.Decode(&cronJob)

		if err != nil {
			log.Println("Failed to decode cron job with error:", err)
			return cronJobs, err
		}

		cronJobs = append(cronJobs, &cronJob)
	}

	return cronJobs, nil
}
