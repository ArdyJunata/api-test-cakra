package views

import (
	"net/http"

	"github.com/ArdyJunata/api-test-cakra/constants"
)

type APIResponse struct {
	Status  int            `json:"status"`
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Payload interface{}    `json:"payload,omitempty"`
	Query   *QueryResponse `json:"query,omitempty"`
}

type QueryResponse struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func SuccessCreatedAPIResponse(payload interface{}) *APIResponse {
	return &APIResponse{
		Status:  http.StatusCreated,
		Success: true,
		Message: constants.SUCCESS_CREATED,
		Payload: payload,
	}
}

func SuccessFindSingleAPIResponse(payload interface{}) *APIResponse {
	return &APIResponse{
		Status:  http.StatusOK,
		Success: true,
		Message: constants.SUCCESS_GET_DATA,
		Payload: payload,
	}
}

func SuccessFindAllAPIResponse(payload interface{}, page, limit, total int) *APIResponse {
	return &APIResponse{
		Status:  http.StatusOK,
		Success: true,
		Message: constants.SUCCESS_GET_DATA,
		Payload: payload,
	}
}

func SuccessUpdateAPIResponse(payload interface{}) *APIResponse {
	return &APIResponse{
		Status:  http.StatusOK,
		Success: true,
		Message: constants.SUCCESS_UPDATE_DATA,
		Payload: payload,
	}
}

func BadRequestAPIResponse(err error, payload interface{}) *APIResponse {
	return buildErrorAPIResponse(http.StatusBadRequest, err.Error(), payload)
}

func NotFoundAPIResponse(err error, payload interface{}) *APIResponse {
	return buildErrorAPIResponse(http.StatusNotFound, err.Error(), payload)
}

func InternalServerErrorAPIResponse(err error, payload interface{}) *APIResponse {
	return buildErrorAPIResponse(http.StatusInternalServerError, err.Error(), payload)
}

func buildErrorAPIResponse(status int, message string, payload interface{}) *APIResponse {
	return &APIResponse{
		Status:  status,
		Success: false,
		Message: message,
		Payload: payload,
	}
}
