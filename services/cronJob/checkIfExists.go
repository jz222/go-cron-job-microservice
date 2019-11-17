package cronjob

import (
	"context"
	"go-cron-job-microservice/libs/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CheckIfExists checks is a cron job with the given ID exists in the database.
func CheckIfExists(jobID primitive.ObjectID) (bool, error) {
	collection := mongodb.GetClient().Collection("cronjob")

	limit := int64(1)

	exists, err := collection.CountDocuments(context.TODO(), bson.M{"_id": jobID}, &options.CountOptions{Limit: &limit})
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}
