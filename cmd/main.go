package main

import (
	"fmt"

	"github.com/ArdyJunata/api-test-cakra/config"
	"github.com/ArdyJunata/api-test-cakra/database"
	"github.com/ArdyJunata/api-test-cakra/middleware"
	"github.com/ArdyJunata/api-test-cakra/pkg/logger"
	"github.com/ArdyJunata/api-test-cakra/server"
	"github.com/ArdyJunata/api-test-cakra/server/controllers"
	"github.com/ArdyJunata/api-test-cakra/server/repositories/gorm"
	carSvc "github.com/ArdyJunata/api-test-cakra/server/service/car"
	clubSvc "github.com/ArdyJunata/api-test-cakra/server/service/club"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log := logger.NewLog()

	config.InitConfig("config/static.env")

	httprouter := httprouter.New()

	db := database.ConnectPostgres()
	if db == nil {
		panic("cant connect to database")
	}

	middleware := middleware.NewMiddleware(log)

	carRepo := gorm.NewCarRepo(db)
	carService := carSvc.NewCarService(carRepo, log)
	carController := controllers.NewCarController(carService)

	clubRepo := gorm.NewClubRepo(db)
	clubService := clubSvc.NewClubService(clubRepo, log)
	clubController := controllers.NewClubController(clubService)

	letterController := controllers.NewLetterController()

	baseController := controllers.NewController()
	baseController = baseController.SetCarController(carController)
	baseController = baseController.SetClubController(clubController)
	baseController = baseController.SetLetterController(letterController)

	router := server.NewRouter(fmt.Sprintf(":%s", config.GetString(config.APP_PORT)), log, baseController, httprouter, middleware)

	router.StartServer()
}
