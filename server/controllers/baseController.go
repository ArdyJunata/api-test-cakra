package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArdyJunata/api-test-cakra/pkg/logger"
	"github.com/ArdyJunata/api-test-cakra/server/views"
)

type Controller struct {
	Car    *CarController
	Club   *ClubController
	Letter *LetterController
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) SetCarController(Car *CarController) *Controller {
	c.Car = Car
	return c
}

func (c *Controller) SetClubController(Club *ClubController) *Controller {
	c.Club = Club
	return c
}

func (c *Controller) SetLetterController(Letter *LetterController) *Controller {
	c.Letter = Letter
	return c
}

func WriteJsonResponse(ctx context.Context, rw http.ResponseWriter, payload *views.APIResponse) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("X-Tracer-ID", fmt.Sprintf("%s", ctx.Value(logger.TRACER_ID)))
	rw.WriteHeader(payload.Status)
	json.NewEncoder(rw).Encode(payload)
}
