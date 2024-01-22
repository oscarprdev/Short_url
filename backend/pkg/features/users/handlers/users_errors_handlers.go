package users

import (
	"net/http"
	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
)

func handleUserErrors(w http.ResponseWriter, err error) error {
	switch e := err.(type) {
	case *errors.BadRequestError:
		status := 400
		errorResponse := api.GetUsers400JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitGetUsersResponse(w)
	case *errors.UnauthorizedError:
		status := 401
		errorResponse := api.GetUsers401JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitGetUsersResponse(w)
	default:
		return errors.DefaultErrorResponse(w)
	}
}
