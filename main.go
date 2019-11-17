package main

import (
	"go-cron-job-microservice/initialization"
	"go-cron-job-microservice/server"
	cronjobmanager "go-cron-job-microservice/services/cronJobManager"
)

func init() {
	initialization.InitEnv()
	initialization.InitDatabase()
}

func main() {
	c := cronjobmanager.GetManager()
	c.Start()

	server.Start()
}
