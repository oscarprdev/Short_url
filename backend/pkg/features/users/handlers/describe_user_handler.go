package users

import (
	"fmt"
	api "short_url/pkg/api"
	userUc "short_url/pkg/features/users/usecases"

	"github.com/labstack/echo/v4"
)

func HandlerDescribeUser(userUc *userUc.UserUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		w := c.Response().Writer

		authToken := c.Request().Header.Get("Authorization")
		idQueryParam := c.QueryParam("id")

		fmt.Println("HELO")

		describeResponse, err := userUc.DescribeUsers(ctx, authToken, idQueryParam)
		if err != nil {
			return handleUserErrors(w, err)
		}

		status := int(200)
		successResponse := api.GetUsersDescribe200JSONResponse{
			Status: &status,
			User:   describeResponse.User,
			Urls:   describeResponse.Urls,
		}

		return successResponse.VisitGetUsersDescribeResponse(w)
	}
}
