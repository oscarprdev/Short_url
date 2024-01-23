package handlers

import (
	"net/http"
	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
)

func handleUrlsErrors(w http.ResponseWriter, err error) error {
	switch e := err.(type) {
	case *errors.BadRequestError:
		status := 400
		errorResponse := api.PostUrl400JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitPostUrlResponse(w)
	case *errors.UnauthorizedError:
		status := 401
		errorResponse := api.PostUrl401JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitPostUrlResponse(w)
	case *errors.InternalError:
		status := 500
		errorResponse := api.PostUrl500JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitPostUrlResponse(w)
	default:
		return errors.DefaultErrorResponse(w)
	}
}
