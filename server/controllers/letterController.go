package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ArdyJunata/api-test-cakra/pkg/logger"
	"github.com/ArdyJunata/api-test-cakra/server/params"
	"github.com/ArdyJunata/api-test-cakra/server/views"
	"github.com/julienschmidt/httprouter"
)

type LetterController struct {
}

func NewLetterController() *LetterController {
	return &LetterController{}
}

func (l *LetterController) IsContainLetters(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req params.Letter
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		bad := views.BadRequestAPIResponse(err, err.Error())
		WriteJsonResponse(r.Context(), w, bad)
		return
	}

	contains := strings.Contains(req.FirstWord, req.LastWord)

	fmt.Println(req.FirstWord)
	fmt.Println(req.LastWord)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Tracer-ID", fmt.Sprintf("%s", r.Context().Value(logger.TRACER_ID)))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contains)

}
