package service

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"wb-report-downloader/internal/report"
	reportdb "wb-report-downloader/internal/report/db"
	"wb-report-downloader/internal/seller"
	sellerdb "wb-report-downloader/internal/seller/db"
	"wb-report-downloader/internal/task"
	taskdb "wb-report-downloader/internal/task/db"
	"wb-report-downloader/internal/wb_request"
	"wb-report-downloader/pkg/client/postgresql"
	"wb-report-downloader/pkg/slice"
)

func Serve(client postgresql.Client, sleepIntervalSec uint) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	tiker := time.NewTicker(time.Duration(sleepIntervalSec) * time.Second)

	for {
		downloadReports(client)
		downloadDetailedReports(client)
		log.Printf("Sleep for %v secs\n", sleepIntervalSec)
		select {
		case <-tiker.C:
			continue
		case <-ctx.Done():
			log.Printf("Shutting down...\n")
			return
		}
	}

}

func downloadReports(postgres postgresql.Client) {
	log.Printf("Getting sellers...\n")
	sellerRepository := sellerdb.NewRepository(postgres)
	sellers, err := sellerRepository.GetSellers(context.TODO())
	if err != nil {
		log.Fatalf("[err] GetSellers: %v\n", err)
		return
	}
	var sellersString string
	for _, seller := range sellers {
		sellersString += fmt.Sprintf("[%d] %v, ", seller.ID, seller.Name)
	}
	log.Printf("Sellers: %v\n", sellersString)

	for _, seller := range sellers {
		handleSeller(seller, postgres)
	}
}

func downloadDetailedReports(postgres postgresql.Client) {
	const (
		kTasksLimitPerIteration = 2
	)

	var done = make(chan bool)

	go func() {
		for {
			log.Printf("Searching for report download task...\n")
			taskRepository := taskdb.NewRepository(postgres)
			tasks, err := taskRepository.GetDownloadTasks(context.TODO(), kTasksLimitPerIteration)
			if err != nil {
				log.Fatalf("GetDownloadTask: %v\n", err)
			}
			if len(tasks) <= 0 {
				done <- true
			}

			for _, t := range tasks {
				log.Printf("Handle task. report_id: %v seller_id: %v \n", t.ReportID, t.SellerID)
				task.Handle(t, taskRepository, postgres)
			}
		}
	}()

	<-done
	log.Printf("All tasks done\n")
}

func handleSeller(seller seller.Seller, postgres postgresql.Client) {
	log.Printf("[%d] Handling seller: [%d] %v\n", seller.ID, seller.ID, seller.Name)
	log.Printf("[%d] Sending http request 'report'... \n", seller.ID)
	r := wb_request.GetReports(seller.XSupplierID, seller.WBToken)
	log.Printf("[%d] Http request ... OK \n", seller.ID)

	allReportsIds := report.GetReportIds(r)
	reportRepository := reportdb.NewRepository(postgres)
	foundIds, err := reportRepository.FindAll(context.TODO(), allReportsIds)
	if err != nil {
		log.Printf("[%d][err] report FindAll: %v\n", seller.ID, err)
		return
	}

	newReportsIds := slice.Difference(allReportsIds, foundIds)

	log.Printf("[%d] newReports: %v\n", seller.ID, newReportsIds)

	reportsForSave := report.GetReportsByIds(r, newReportsIds)

	for _, newReport := range reportsForSave {
		log.Printf("[%d] Saving report: %d ...\n", seller.ID, newReport.Id)
		err := reportRepository.Save(context.TODO(), seller.ID, &newReport)
		if err != nil {
			log.Printf("[%d] Error save report: %d", seller.ID, err)
		} else {
			log.Printf("[%d] Saving report: %d ... OK\n", seller.ID, newReport.Id)
		}
	}

	log.Printf("[%d] Handling seller: [%d] %v ... OK\n", seller.ID, seller.ID, seller.Name)
}
