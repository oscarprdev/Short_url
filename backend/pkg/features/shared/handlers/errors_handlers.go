package errors

import (
	"encoding/json"
	"net/http"
	api "short_url/pkg/api"
)

type BadRequestError struct {
	Details string
}

func (e *BadRequestError) Error() string {
	return e.Details
}

type UnauthorizedError struct {
	Details string
}

func (e *UnauthorizedError) Error() string {
	return e.Details
}

func DefaultErrorResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	details := "Error not found"
	status := 404

	errorResponse := api.Error{
		Details: &details,
		Status:  &status,
	}

	return json.NewEncoder(w).Encode(errorResponse)
}
