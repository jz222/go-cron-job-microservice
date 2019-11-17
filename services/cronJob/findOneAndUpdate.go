package cronjob

import (
	"context"
	"go-cron-job-microservice/libs/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

// FindOneAndUpdate finds and updates a document with the given filter and update.
func FindOneAndUpdate(filter bson.M, update bson.M) error {
	collection := mongodb.GetClient().Collection("cronjob")

	res := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}
