package repositories

import (
	"context"

	"github.com/ArdyJunata/api-test-cakra/server/models"
)

type CarRepo interface {
	CreateCar(ctx context.Context, car *models.Car) error
	GetGroupedBrandCars(ctx context.Context) (*[]models.Car, error)
}

type ClubRepo interface {
	CreateClub(ctx context.Context, club *models.Club) error
	LeagueStandings(ctx context.Context) (*[]models.Club, error)
	ClubStandings(ctx context.Context, clubname string) (*models.Club, error)
	RecordGame(ctx context.Context, clubhomename, clubawayname, winner *models.Club, score int) error
}
