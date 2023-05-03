package main

import (
	"context"
	"log"
	"time"
	"wb-report-downloader/internal/config"
	"wb-report-downloader/internal/task"
	"wb-report-downloader/internal/task/db"
	"wb-report-downloader/pkg/client/postgresql"
)

func main() {
	log.Printf("Detailed report downloader\n")

	log.Printf("Reading config...")
	c, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("LoadConfig: %v\n", err)
	}
	log.Printf("Reading config... OK")

	log.Printf("Connecting database...\n")
	postgreSQLClient, err := postgresql.NewClient(context.TODO(),
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	log.Printf("Connecting database... OK\n")

	for {
		if(!workIteration(postgreSQLClient)) {
			log.Printf("New tasks not found. Sleep for %v secs", c.SleepOnTaskNotFoundSec)
			time.Sleep(time.Duration(c.SleepOnTaskNotFoundSec) * time.Second)
		}
	}
}

func workIteration(db_client postgresql.Client) bool {
	const (
		kTasksLimitPerIteration = 2
	)

	log.Printf("Searching for report download task...\n")
	taskRepository := taskdb.NewRepository(db_client)
	tasks, err := taskRepository.GetDownloadTasks(context.TODO(), kTasksLimitPerIteration)
	if err != nil {
		log.Fatalf("GetDownloadTask: %v\n", err)
	}
	if len(tasks) <= 0 {
		return false
	}

	for _, t := range tasks {
		log.Printf("Handle task. report_id: %v seller_id: %v \n", t.ReportID, t.SellerID)
		task.Handle(t, taskRepository, db_client)
	}
	return true
}

