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
		errorResponse := api.GetUsersList400JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitGetUsersListResponse(w)
	case *errors.UnauthorizedError:
		status := 401
		errorResponse := api.GetUsersList401JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitGetUsersListResponse(w)
	case *errors.InternalError:
		status := 401
		errorResponse := api.GetUsersList500JSONResponse{
			Details: &e.Details,
			Status:  &status,
		}

		return errorResponse.VisitGetUsersListResponse(w)
	default:
		return errors.DefaultErrorResponse(w)
	}
}
