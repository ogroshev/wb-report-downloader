package main

import (
	"log"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"wb-report-downloader/internal/config"
	"wb-report-downloader/internal/service"
	"wb-report-downloader/pkg/client/postgresql"

)

func main()  {
	postgreSQLClient := start()
	service.Serve(postgreSQLClient)
}

func start() *pgxpool.Pool {
	log.Printf("Report downloader\n")

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

	return postgreSQLClient
}