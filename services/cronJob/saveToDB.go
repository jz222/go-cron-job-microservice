package cronjob

import (
	"context"
	"go-cron-job-microservice/libs/mongodb"
	"go-cron-job-microservice/models"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveToDatabase saves a given cron job in the database.
func SaveToDatabase(newCronJob models.CronJob) (primitive.ObjectID, error) {
	collection := mongodb.GetClient().Collection("cronjob")
	result, err := collection.InsertOne(context.TODO(), newCronJob)

	if err != nil {
		log.Println("Failed to save new cron job to database with error:", err)
		return primitive.ObjectID{}, errors.New("An error occured while saving a cron job to the database")
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return objectID, nil
	}

	return primitive.ObjectID{}, nil
}
