package gorm

import (
	"context"
	"errors"

	"github.com/ArdyJunata/api-test-cakra/server/models"
	"github.com/ArdyJunata/api-test-cakra/server/repositories"
	"gorm.io/gorm"
)

type carRepo struct {
	db *gorm.DB
}

func NewCarRepo(db *gorm.DB) repositories.CarRepo {
	return &carRepo{
		db: db,
	}
}

func (c *carRepo) GetGroupedBrandCars(ctx context.Context) (*[]models.Car, error) {
	var cars []models.Car

	res := c.db.Raw("SELECT brand from cars group by brand").Scan(&cars)
	if res.Error != nil {
		return nil, res.Error
	}

	if errors.Is(res.Error, gorm.ErrRecordNotFound) || res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &cars, nil
}

func (c *carRepo) CreateCar(ctx context.Context, car *models.Car) error {
	return c.db.Create(car).Error
}
