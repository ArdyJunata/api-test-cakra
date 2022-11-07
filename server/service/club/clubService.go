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

type ClubService struct {
	clubRepo repositories.ClubRepo
	log      logger.Logger
}

func NewClubService(clubRepo repositories.ClubRepo, log logger.Logger) *ClubService {
	return &ClubService{
		clubRepo: clubRepo,
		log:      log,
	}
}

func makeClubModelFromRequestCreate(req *params.Club) *models.Club {
	return &models.Club{
		ClubName: req.ClubName,
		Point:    req.Point,
	}
}

func (c *ClubService) CreateClub(ctx context.Context, req *params.Club) *views.APIResponse {
	car := makeClubModelFromRequestCreate(req)

	err := c.clubRepo.CreateClub(ctx, car)

	if err != nil {
		c.log.Errorf(ctx, "error when try to create new club stack with error %s", err.Error())
		return views.InternalServerErrorAPIResponse(err, nil)
	}

	return views.SuccessCreatedAPIResponse(nil)
}

func (c *ClubService) GetAllClub(ctx context.Context) *views.APIResponse {
	clubs, err := c.clubRepo.LeagueStandings(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.NotFoundAPIResponse(err, nil)
		}
		return views.InternalServerErrorAPIResponse(err, nil)
	}

	return views.SuccessFindAllAPIResponse(clubs, 0, 0, len(*clubs))
}

func (c *ClubService) ClubStandings(ctx context.Context, clubname string) *views.APIResponse {
	club, err := c.clubRepo.ClubStandings(ctx, clubname)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.NotFoundAPIResponse(err, nil)
		}
		return views.InternalServerErrorAPIResponse(err, nil)
	}

	clubs, err := c.clubRepo.LeagueStandings(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.NotFoundAPIResponse(err, nil)
		}
		return views.InternalServerErrorAPIResponse(err, nil)
	}

	rank := 0

	var req params.ClubStandings

	for i, e := range *clubs {
		if club.ClubName == e.ClubName {
			rank = i + 1
		}
	}

	req.ClubName = club.ClubName
	req.Standing = rank

	return views.SuccessFindSingleAPIResponse(req)

}

func (c *ClubService) RecordGame(ctx context.Context, clubhomename, clubawayname, winner string, score int) *views.APIResponse {
	if score == 1 {
		var win *models.Club
		home, err := c.clubRepo.ClubStandings(ctx, clubhomename)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return views.NotFoundAPIResponse(err, nil)
			}
			return views.InternalServerErrorAPIResponse(err, nil)
		}

		away, err := c.clubRepo.ClubStandings(ctx, clubawayname)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return views.NotFoundAPIResponse(err, nil)
			}
			return views.InternalServerErrorAPIResponse(err, nil)
		}

		err = c.clubRepo.RecordGame(ctx, home, away, win, score)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return views.NotFoundAPIResponse(err, nil)
			}
			return views.InternalServerErrorAPIResponse(err, nil)
		}
	} else {
		var home *models.Club
		var away *models.Club
		win, err := c.clubRepo.ClubStandings(ctx, winner)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return views.NotFoundAPIResponse(err, nil)
			}
			return views.InternalServerErrorAPIResponse(err, nil)
		}

		err = c.clubRepo.RecordGame(ctx, home, away, win, score)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return views.NotFoundAPIResponse(err, nil)
			}
			return views.InternalServerErrorAPIResponse(err, nil)
		}
	}

	return views.SuccessUpdateAPIResponse(nil)

}
