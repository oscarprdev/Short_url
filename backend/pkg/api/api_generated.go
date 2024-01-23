// Package short_url provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package short_url

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Error defines model for Error.
type Error struct {
	// Details Error details
	Details *string `json:"details,omitempty"`

	// Status Response code
	Status *int `json:"status,omitempty"`
}

// Url defines model for Url.
type Url struct {
	// CreatedAt The timestamp when the URL was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// ExpiresAt Expiration date and time for the URL.
	ExpiresAt *time.Time `json:"expiresAt"`

	// Id The unique identifier for the URL.
	Id *openapi_types.UUID `json:"id,omitempty"`

	// OriginalUrl The original URL.
	OriginalUrl *string `json:"originalUrl,omitempty"`

	// ShortUrl The short URL.
	ShortUrl *string `json:"shortUrl,omitempty"`

	// TitleUrl The URL title.
	TitleUrl *string `json:"titleUrl,omitempty"`

	// UpdatedAt The timestamp when the URL was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// User defines model for User.
type User struct {
	// CreatedAt The timestamp when the user was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Email The user's email.
	Email *openapi_types.Email `json:"email,omitempty"`

	// Id The unique identifier for the user.
	Id *string `json:"id,omitempty"`

	// Name The user's first name.
	Name *string `json:"name,omitempty"`

	// Picture The user's picture URL.
	Picture *string `json:"picture,omitempty"`

	// UpdatedAt The timestamp when the user was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// PostUrlJSONBody defines parameters for PostUrl.
type PostUrlJSONBody struct {
	// OriginalUrl The original URL.
	OriginalUrl *string `json:"originalUrl,omitempty"`
}

// PostUrlParams defines parameters for PostUrl.
type PostUrlParams struct {
	// UserId User identifier
	UserId *string `form:"UserId,omitempty" json:"UserId,omitempty"`

	// Authorization Auth token for user authorization
	Authorization *string `json:"Authorization,omitempty"`
}

// PostUrlJSONRequestBody defines body for PostUrl for application/json ContentType.
type PostUrlJSONRequestBody PostUrlJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /auth)
	GetAuth(ctx echo.Context) error

	// (GET /auth/logout)
	GetAuthLogout(ctx echo.Context) error

	// (POST /url)
	PostUrl(ctx echo.Context, params PostUrlParams) error

	// (GET /users)
	GetUsers(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAuth converts echo context to params.
func (w *ServerInterfaceWrapper) GetAuth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAuth(ctx)
	return err
}

// GetAuthLogout converts echo context to params.
func (w *ServerInterfaceWrapper) GetAuthLogout(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAuthLogout(ctx)
	return err
}

// PostUrl converts echo context to params.
func (w *ServerInterfaceWrapper) PostUrl(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostUrlParams
	// ------------- Optional query parameter "UserId" -------------

	err = runtime.BindQueryParameter("form", true, false, "UserId", ctx.QueryParams(), &params.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UserId: %s", err))
	}

	headers := ctx.Request().Header
	// ------------- Optional header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for Authorization, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, valueList[0], &Authorization)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Authorization: %s", err))
		}

		params.Authorization = &Authorization
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUrl(ctx, params)
	return err
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/auth", wrapper.GetAuth)
	router.GET(baseURL+"/auth/logout", wrapper.GetAuthLogout)
	router.POST(baseURL+"/url", wrapper.PostUrl)
	router.GET(baseURL+"/users", wrapper.GetUsers)

}

type GetAuthRequestObject struct {
}

type GetAuthResponseObject interface {
	VisitGetAuthResponse(w http.ResponseWriter) error
}

type GetAuth200JSONResponse struct {
	AuthInfo *User `json:"authInfo,omitempty"`
	Status   *int  `json:"status,omitempty"`
}

func (response GetAuth200JSONResponse) VisitGetAuthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetAuth400JSONResponse Error

func (response GetAuth400JSONResponse) VisitGetAuthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetAuth401JSONResponse Error

func (response GetAuth401JSONResponse) VisitGetAuthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetAuth500JSONResponse Error

func (response GetAuth500JSONResponse) VisitGetAuthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetAuthLogoutRequestObject struct {
}

type GetAuthLogoutResponseObject interface {
	VisitGetAuthLogoutResponse(w http.ResponseWriter) error
}

type GetAuthLogout200JSONResponse struct {
	Status *int `json:"status,omitempty"`
}

func (response GetAuthLogout200JSONResponse) VisitGetAuthLogoutResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetAuthLogout400JSONResponse Error

func (response GetAuthLogout400JSONResponse) VisitGetAuthLogoutResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetAuthLogout401JSONResponse Error

func (response GetAuthLogout401JSONResponse) VisitGetAuthLogoutResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetAuthLogout500JSONResponse Error

func (response GetAuthLogout500JSONResponse) VisitGetAuthLogoutResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostUrlRequestObject struct {
	Params PostUrlParams
	Body   *PostUrlJSONRequestBody
}

type PostUrlResponseObject interface {
	VisitPostUrlResponse(w http.ResponseWriter) error
}

type PostUrl200JSONResponse struct {
	Status *int `json:"status,omitempty"`
	Url    *Url `json:"url,omitempty"`
}

func (response PostUrl200JSONResponse) VisitPostUrlResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUrl400JSONResponse Error

func (response PostUrl400JSONResponse) VisitPostUrlResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostUrl401JSONResponse Error

func (response PostUrl401JSONResponse) VisitPostUrlResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type PostUrl500JSONResponse Error

func (response PostUrl500JSONResponse) VisitPostUrlResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse struct {
	Status *int    `json:"status,omitempty"`
	Users  *[]User `json:"users,omitempty"`
}

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsers400JSONResponse Error

func (response GetUsers400JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetUsers401JSONResponse Error

func (response GetUsers401JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetUsers500JSONResponse Error

func (response GetUsers500JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /auth)
	GetAuth(ctx context.Context, request GetAuthRequestObject) (GetAuthResponseObject, error)

	// (GET /auth/logout)
	GetAuthLogout(ctx context.Context, request GetAuthLogoutRequestObject) (GetAuthLogoutResponseObject, error)

	// (POST /url)
	PostUrl(ctx context.Context, request PostUrlRequestObject) (PostUrlResponseObject, error)

	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetAuth operation middleware
func (sh *strictHandler) GetAuth(ctx echo.Context) error {
	var request GetAuthRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAuth(ctx.Request().Context(), request.(GetAuthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAuth")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetAuthResponseObject); ok {
		return validResponse.VisitGetAuthResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetAuthLogout operation middleware
func (sh *strictHandler) GetAuthLogout(ctx echo.Context) error {
	var request GetAuthLogoutRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAuthLogout(ctx.Request().Context(), request.(GetAuthLogoutRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAuthLogout")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetAuthLogoutResponseObject); ok {
		return validResponse.VisitGetAuthLogoutResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUrl operation middleware
func (sh *strictHandler) PostUrl(ctx echo.Context, params PostUrlParams) error {
	var request PostUrlRequestObject

	request.Params = params

	var body PostUrlJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUrl(ctx.Request().Context(), request.(PostUrlRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUrl")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUrlResponseObject); ok {
		return validResponse.VisitPostUrlResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
