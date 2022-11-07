package service

import (
	"context"

	"github.com/ArdyJunata/api-test-cakra/pkg/logger"
	"github.com/ArdyJunata/api-test-cakra/server/models"
	"github.com/ArdyJunata/api-test-cakra/server/params"
	"github.com/ArdyJunata/api-test-cakra/server/repositories"
	"github.com/ArdyJunata/api-test-cakra/server/views"
	"gorm.io/gorm"
)

type CarService struct {
	carRepo repositories.CarRepo
	log     logger.Logger
}

func NewCarService(carRepo repositories.CarRepo, log logger.Logger) *CarService {
	return &CarService{
		carRepo: carRepo,
		log:     log,
	}
}

func makeCarModelFromRequestCreate(req *params.Car) *models.Car {
	return &models.Car{
		Price: req.Price,
		Brand: req.Brand,
		Type:  req.Type,
	}
}

func (c *CarService) GetGroupedBrandCars(ctx context.Context) *views.APIResponse {
	cars, err := c.carRepo.GetGroupedBrandCars(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.NotFoundAPIResponse(err, nil)
		}
		return views.InternalServerErrorAPIResponse(err, nil)
	}

	return views.SuccessFindAllAPIResponse(cars, 0, 0, len(*cars))
}

func (c *CarService) CreateCar(ctx context.Context, req *params.Car) *views.APIResponse {
	car := makeCarModelFromRequestCreate(req)

	err := c.carRepo.CreateCar(ctx, car)

	if err != nil {
		c.log.Errorf(ctx, "error when try to create new car stack with error %s", err.Error())
		return views.InternalServerErrorAPIResponse(err, nil)
	}

	return views.SuccessCreatedAPIResponse(nil)
}
