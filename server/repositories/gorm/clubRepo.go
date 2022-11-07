package gorm

import (
	"context"
	"fmt"

	"github.com/ArdyJunata/api-test-cakra/server/models"
	"github.com/ArdyJunata/api-test-cakra/server/repositories"
	"gorm.io/gorm"
)

type clubRepo struct {
	db *gorm.DB
}

func NewClubRepo(db *gorm.DB) repositories.ClubRepo {
	return &clubRepo{
		db: db,
	}
}

func (c *clubRepo) CreateClub(ctx context.Context, club *models.Club) error {
	return c.db.Create(club).Error
}

func (c *clubRepo) LeagueStandings(ctx context.Context) (*[]models.Club, error) {
	var club []models.Club

	err := c.db.Order("point desc").Find(&club).Error

	if err != nil {
		return nil, err
	}

	return &club, nil
}

func (c *clubRepo) ClubStandings(ctx context.Context, clubname string) (*models.Club, error) {
	var club models.Club

	err := c.db.Where("club_name = ?", clubname).Find(&club).Error

	if err != nil {
		return nil, err
	}

	return &club, nil
}

func (c *clubRepo) RecordGame(ctx context.Context, clubhome, clubaway, winner *models.Club, score int) error {
	if score == 1 {
		final_point := score + int(clubhome.Point)
		clubhome.Point = uint64(final_point)
		err := c.db.Where("club_name=?", clubhome.ClubName).Updates(clubhome).Error
		if err != nil {
			return err
		}
		final_point = score + int(clubaway.Point)
		clubaway.Point = uint64(final_point)
		err = c.db.Where("club_name=?", clubaway.ClubName).Updates(clubaway).Error
		if err != nil {
			return err
		}
	} else {
		final_point := score + int(winner.Point)
		// err := c.db.Raw("UPDATE clubs SET point=? WHERE club_name=?", final_point, winner.ClubName).Error
		winner.Point = uint64(final_point)
		err := c.db.Where("club_name=?", winner.ClubName).Updates(winner).Error
		if err != nil {
			return err
		}
		fmt.Println(winner)
	}
	return nil
}
