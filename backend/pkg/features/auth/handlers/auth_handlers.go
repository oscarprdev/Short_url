package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	authUc "short_url/pkg/features/auth/usecases"
	errors "short_url/pkg/features/shared/handlers"

	"github.com/labstack/echo/v4"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func provideHttpResponse(e *echo.Response) *http.Response {
	return &http.Response{}
}

func handleAuthUserUsecase(ctx context.Context, w http.ResponseWriter, authUc *authUc.AuthUsecases, gu *goth.User) error {
	err := authUc.AuthUser(ctx, gu)
	if err != nil {
		newError := &errors.InternalError{
			Details: fmt.Sprintf("Internal error: %v", err),
		}
		handleUserErrors(w, newError)
	}

	return nil
}

func HandlerAuthCallback(authUc *authUc.AuthUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		r := c.Request()
		ctx := c.Request().Context()

		gothUser, err := gothic.CompleteUserAuth(w, r)

		if err != nil {
			newError := &errors.UnauthorizedError{
				Details: fmt.Sprintf("Error authorizing user: %v", err),
			}
			return handleUserErrors(w, newError)
		}

		err = handleAuthUserUsecase(ctx, w, authUc, &gothUser)
		if err != nil {
			newError := &errors.InternalError{
				Details: fmt.Sprintf("Error providing user from db: %v", err),
			}
			return handleUserErrors(w, newError)
		}

		clientUrl := os.Getenv("CLIENT_URL")
		// Redirect to a success page or another route
		path := clientUrl + "user/" + gothUser.UserID
		http.Redirect(w, r, path, http.StatusFound)
		return nil
	}
}

func HandlerAuth(authUc *authUc.AuthUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		r := c.Request()
		ctx := c.Request().Context()

		gothUser, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			gothic.BeginAuthHandler(w, r)
			return nil
		}

		err = handleAuthUserUsecase(ctx, w, authUc, &gothUser)
		if err != nil {
			newError := &errors.InternalError{
				Details: fmt.Sprintf("Error providing user from db: %v", err),
			}
			return handleUserErrors(w, newError)
		}

		clientUrl := os.Getenv("CLIENT_URL")
		path := clientUrl + "user/" + gothUser.UserID
		http.Redirect(w, r, path, http.StatusFound)
		return nil
	}
}

func HandlerLogout(authUc *authUc.AuthUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		r := c.Request()

		gothic.Logout(w, r)

		clientUrl := os.Getenv("CLIENT_URL")
		http.Redirect(w, r, clientUrl, http.StatusFound)
		return nil
	}
}
