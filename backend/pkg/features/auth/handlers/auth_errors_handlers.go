package handlers

import (
	"net/http"
	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
)

func handleUserErrors(w http.ResponseWriter, err error) error {
	switch e := err.(type) {
	case *errors.BadRequestError:
		status := 400
		errorResponse := api.GetAuth400JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitGetAuthResponse(w)
	case *errors.UnauthorizedError:
		status := 401
		errorResponse := api.GetAuth401JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitGetAuthResponse(w)
	case *errors.InternalError:
		status := 500
		errorResponse := api.GetAuth500JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitGetAuthResponse(w)
	default:
		return errors.DefaultErrorResponse(w)
	}
}
