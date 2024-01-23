package handlers

import (
	"fmt"
	"net/http"
	userUc "short_url/pkg/features/auth/usecases"
	errors "short_url/pkg/features/shared/handlers"

	"github.com/labstack/echo/v4"

	"github.com/markbates/goth/gothic"
)

func provideHttpResponse(e *echo.Response) *http.Response {
	return &http.Response{}
}

func HandlerAuthCallback(userUc *userUc.AuthUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		r := c.Request()

		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			newError := &errors.UnauthorizedError{
				Details: fmt.Sprintf("Error authorizing user: %v", err),
			}
			return handleUserErrors(w, newError)
		}

		fmt.Println(user)

		// Redirect to a success page or another route
		http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
		return nil
	}
}

func HandlerAuth(userUc *userUc.AuthUsecases) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		r := c.Request()

		gothUser, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			gothic.BeginAuthHandler(w, r)
		}

		fmt.Println("Success")
		fmt.Println(gothUser)

		http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
		return nil
	}
}
