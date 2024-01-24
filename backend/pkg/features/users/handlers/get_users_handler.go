package users

import (
	api "short_url/pkg/api"
	userUc "short_url/pkg/features/users/usecases"

	"github.com/labstack/echo/v4"
)

func HandlerGetUsers(userUc *userUc.UserUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		w := c.Response().Writer

		users, err := userUc.Repo.GetUsers(ctx)
		if err != nil {
			return handleUserErrors(w, err)
		}

		status := int(200)
		successResponse := api.GetUsersList200JSONResponse{
			Status: &status,
			Users:  users,
		}

		return successResponse.VisitGetUsersListResponse(w)
	}
}
