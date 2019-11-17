package initialization

import (
	"go-cron-job-microservice/keys"
	"go-cron-job-microservice/libs/mongodb"
	"log"
)

// InitEnv initializes all environment variables.
func InitEnv() {
	keys.GetKeys()

	log.Println("✅ environment variables initialized successfully")
}

// InitDatabase initializes the MongoDB.
func InitDatabase() {
	mongodb.InitiateDatabase()
}
