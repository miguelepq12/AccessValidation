package main

import (
	"AccessValidation/app/application"
	"AccessValidation/app/infrastructure/kafka/access/producer"
	mysql "AccessValidation/app/infrastructure/mysql"
	"AccessValidation/app/infrastructure/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.HideBanner = true

	db,_ := mysql.NewMysqlConnection()
	defer db.Close()
	repository := mysql.NewAccessMysqlRepository(db)
	producerAccess := producer.NewAccessRecordKafkaProducer()
	defer producerAccess.Close()
	accessUseCase := application.NewAccessUseCase(repository,producerAccess)
	accessHandler := web.NewAccessHandler(accessUseCase)

	e.POST("/validate", accessHandler.ValidateAccess)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

