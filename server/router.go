package server

import (
	"context"
	"net/http"

	"github.com/ArdyJunata/api-test-cakra/middleware"
	"github.com/ArdyJunata/api-test-cakra/pkg/logger"
	"github.com/ArdyJunata/api-test-cakra/server/controllers"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router         *httprouter.Router
	appPort        string
	log            logger.Logger
	baseController *controllers.Controller
	middleware     *middleware.Middleware
}

func NewRouter(appPort string, log logger.Logger, baseController *controllers.Controller, router *httprouter.Router, middleware *middleware.Middleware) *Router {
	return &Router{
		router:         router,
		appPort:        appPort,
		log:            log,
		baseController: baseController,
		middleware:     middleware,
	}
}

func (r *Router) StartServer() {
	r.buildRoute()

	r.log.Infof(context.Background(), "server running at port %s", r.appPort)
	http.ListenAndServe(r.appPort, r.router)
}

func (r *Router) buildRoute() {
	r.carRouter()
	r.clubRouter()
	r.containLetter()
}

func (r *Router) carRouter() {
	r.router.GET("/cars", r.middleware.Tracer(r.baseController.Car.GetGroupedBrandCars))
	r.router.POST("/cars", r.middleware.Tracer(r.baseController.Car.CreateCar))
}

func (r *Router) clubRouter() {
	r.router.POST("/club", r.middleware.Tracer(r.baseController.Club.CreateClub))
	r.router.GET("/club", r.middleware.Tracer(r.baseController.Club.GetAllClub))
	r.router.GET("/club/rank", r.middleware.Tracer(r.baseController.Club.ClubStandings))
	r.router.POST("/club/recordgame", r.middleware.Tracer(r.baseController.Club.RecordGame))
}

func (r *Router) containLetter() {
	r.router.POST("/is-contain-letters", r.middleware.Tracer(r.baseController.Letter.IsContainLetters))
}
