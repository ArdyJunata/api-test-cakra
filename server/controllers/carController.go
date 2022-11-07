package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArdyJunata/api-test-cakra/server/params"
	service "github.com/ArdyJunata/api-test-cakra/server/service/car"
	"github.com/ArdyJunata/api-test-cakra/server/views"
	"github.com/julienschmidt/httprouter"
)

type CarController struct {
	svc *service.CarService
}

func NewCarController(svc *service.CarService) *CarController {
	return &CarController{
		svc: svc,
	}
}

func (c *CarController) GetGroupedBrandCars(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := c.svc.GetGroupedBrandCars(r.Context())

	WriteJsonResponse(r.Context(), w, res)

}

func (c *CarController) CreateCar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req params.Car
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		bad := views.BadRequestAPIResponse(err, err.Error())
		WriteJsonResponse(r.Context(), w, bad)
		return
	}

	res := c.svc.CreateCar(r.Context(), &req)

	WriteJsonResponse(r.Context(), w, res)
}
