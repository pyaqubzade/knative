package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pyaqubzade/knative/client"
	"github.com/pyaqubzade/knative/config"
	"github.com/pyaqubzade/knative/errhandler"
	"github.com/pyaqubzade/knative/handler"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()

	//db.MigrateDB()
	//dbCon := db.ConnectDB()
	//DB, _ := dbCon.DB()
	//defer DB.Close()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          errhandler.ErrorHandler,
	})
	app.Use(recover.New())
	//app.Use(cors.New(cors.Config{
	//	AllowOrigins: "https://avtoyoxla.az",
	//	AllowMethods: "POST",
	//}))
	//app.Use(middleware.NewMDC())
	api := app.Group("/api")

	//reportRepo := db.NewReportRepository(dbCon)
	//actionRepo := db.NewUserActionRepository(dbCon)
	//fmsService := service.NewFmsService(actionRepo)
	//carfaxClient := client.NewCarfaxClient()
	//incidentClient := client.NewClient()
	//smsRadarClient := client.NewSMSRadarClient()
	//pdfConverterClient := client.NewPDFConverterClient()
	//
	//reportService := service.NewReportService(reportRepo, fmsService, carfaxClient, incidentClient, smsRadarClient,
	//	pdfConverterClient)
	c := client.NewClient()
	handler.NewHealthHandler(app)
	handler.NewHandler(api, c)

	port := config.Props.Port
	log.Info("Starting server at port: ", port)
	log.Fatal(app.Listen(":" + port))
}
