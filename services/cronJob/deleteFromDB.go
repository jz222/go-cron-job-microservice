package cronjob

import (
	"context"
	"go-cron-job-microservice/libs/mongodb"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// DeleteFromDB deletes the cron job with the given ID from the database.
func DeleteFromDB(filter bson.M) error {
	collection := mongodb.GetClient().Collection("cronjob")

	if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
		log.Println("Failed to delete cron job from database with error:", err)
		return err
	}

	return nil
}
