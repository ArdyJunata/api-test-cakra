package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/ArdyJunata/api-test-cakra/server/params"
	service "github.com/ArdyJunata/api-test-cakra/server/service/club"
	"github.com/ArdyJunata/api-test-cakra/server/views"
	"github.com/julienschmidt/httprouter"
)

type ClubController struct {
	svc *service.ClubService
}

func NewClubController(svc *service.ClubService) *ClubController {
	return &ClubController{
		svc: svc,
	}
}

func (c *ClubController) CreateClub(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req params.Club
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		bad := views.BadRequestAPIResponse(err, err.Error())
		WriteJsonResponse(r.Context(), w, bad)
		return
	}

	res := c.svc.CreateClub(r.Context(), &req)

	WriteJsonResponse(r.Context(), w, res)
}

func (c *ClubController) GetAllClub(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := c.svc.GetAllClub(r.Context())

	WriteJsonResponse(r.Context(), w, res)

}

func (c *ClubController) ClubStandings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := r.URL.Query()
	clubname := q.Get("clubname")

	res := c.svc.ClubStandings(r.Context(), clubname)

	WriteJsonResponse(r.Context(), w, res)

}

func (c *ClubController) RecordGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req params.RecordGame
	winner := ""
	score := 0
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		bad := views.BadRequestAPIResponse(err, err.Error())
		WriteJsonResponse(r.Context(), w, bad)
		return
	}

	re := regexp.MustCompile("[:\\s]+")
	txt := req.Score

	split := re.Split(txt, -1)

	if split[0] > split[1] {
		winner = req.ClubHomeName
		score = 3
	} else if split[0] < split[1] {
		winner = req.ClubAwayName
		score = 3
	} else if split[0] == split[1] {
		score = 1
	}

	res := c.svc.RecordGame(r.Context(), req.ClubHomeName, req.ClubAwayName, winner, score)

	WriteJsonResponse(r.Context(), w, res)

}
