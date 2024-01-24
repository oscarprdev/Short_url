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

	// Usage The URL usage count.
	Usage *int `json:"usage,omitempty"`
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
	// Id User identifier
	Id *string `form:"Id,omitempty" json:"Id,omitempty"`

	// Authorization Auth token for user authorization
	Authorization *string `json:"Authorization,omitempty"`
}

// PostUrlTitleJSONBody defines parameters for PostUrlTitle.
type PostUrlTitleJSONBody struct {
	// UrlId URL Id.
	UrlId *openapi_types.UUID `json:"urlId,omitempty"`

	// UrlTitle The new URL title.
	UrlTitle *string `json:"urlTitle,omitempty"`
}

// PostUrlTitleParams defines parameters for PostUrlTitle.
type PostUrlTitleParams struct {
	// Id User identifier
	Id *string `form:"Id,omitempty" json:"Id,omitempty"`

	// Authorization Auth token for user authorization
	Authorization *string `json:"Authorization,omitempty"`
}

// PostUrlUseJSONBody defines parameters for PostUrlUse.
type PostUrlUseJSONBody struct {
	// UrlId URL Id.
	UrlId *openapi_types.UUID `json:"urlId,omitempty"`
}

// PostUrlUseParams defines parameters for PostUrlUse.
type PostUrlUseParams struct {
	// Id User identifier
	Id *string `form:"Id,omitempty" json:"Id,omitempty"`

	// Authorization Auth token for user authorization
	Authorization *string `json:"Authorization,omitempty"`
}

// GetUsersDescribeParams defines parameters for GetUsersDescribe.
type GetUsersDescribeParams struct {
	// Id User identifier
	Id *string `form:"Id,omitempty" json:"Id,omitempty"`

	// Authorization Auth token for user authorization
	Authorization *string `json:"Authorization,omitempty"`
}

// PostUrlJSONRequestBody defines body for PostUrl for application/json ContentType.
type PostUrlJSONRequestBody PostUrlJSONBody

// PostUrlTitleJSONRequestBody defines body for PostUrlTitle for application/json ContentType.
type PostUrlTitleJSONRequestBody PostUrlTitleJSONBody

// PostUrlUseJSONRequestBody defines body for PostUrlUse for application/json ContentType.
type PostUrlUseJSONRequestBody PostUrlUseJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /auth)
	GetAuth(ctx echo.Context) error

	// (GET /auth/logout)
	GetAuthLogout(ctx echo.Context) error

	// (POST /url)
	PostUrl(ctx echo.Context, params PostUrlParams) error

	// (POST /url/title)
	PostUrlTitle(ctx echo.Context, params PostUrlTitleParams) error

	// (POST /url/use)
	PostUrlUse(ctx echo.Context, params PostUrlUseParams) error

	// (GET /users/describe)
	GetUsersDescribe(ctx echo.Context, params GetUsersDescribeParams) error

	// (GET /users/list)
	GetUsersList(ctx echo.Context) error
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
	// ------------- Optional query parameter "Id" -------------

	err = runtime.BindQueryParameter("form", true, false, "Id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Id: %s", err))
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

// PostUrlTitle converts echo context to params.
func (w *ServerInterfaceWrapper) PostUrlTitle(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostUrlTitleParams
	// ------------- Optional query parameter "Id" -------------

	err = runtime.BindQueryParameter("form", true, false, "Id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Id: %s", err))
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
	err = w.Handler.PostUrlTitle(ctx, params)
	return err
}

// PostUrlUse converts echo context to params.
func (w *ServerInterfaceWrapper) PostUrlUse(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostUrlUseParams
	// ------------- Optional query parameter "Id" -------------

	err = runtime.BindQueryParameter("form", true, false, "Id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Id: %s", err))
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
	err = w.Handler.PostUrlUse(ctx, params)
	return err
}

// GetUsersDescribe converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersDescribe(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersDescribeParams
	// ------------- Optional query parameter "Id" -------------

	err = runtime.BindQueryParameter("form", true, false, "Id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Id: %s", err))
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
	err = w.Handler.GetUsersDescribe(ctx, params)
	return err
}

// GetUsersList converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersList(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersList(ctx)
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
	router.POST(baseURL+"/url/title", wrapper.PostUrlTitle)
	router.POST(baseURL+"/url/use", wrapper.PostUrlUse)
	router.GET(baseURL+"/users/describe", wrapper.GetUsersDescribe)
	router.GET(baseURL+"/users/list", wrapper.GetUsersList)

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

type PostUrlTitleRequestObject struct {
	Params PostUrlTitleParams
	Body   *PostUrlTitleJSONRequestBody
}

type PostUrlTitleResponseObject interface {
	VisitPostUrlTitleResponse(w http.ResponseWriter) error
}

type PostUrlTitle200JSONResponse struct {
	Status *int `json:"status,omitempty"`
	Url    *Url `json:"url,omitempty"`
}

func (response PostUrlTitle200JSONResponse) VisitPostUrlTitleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUrlTitle400JSONResponse Error

func (response PostUrlTitle400JSONResponse) VisitPostUrlTitleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostUrlTitle401JSONResponse Error

func (response PostUrlTitle401JSONResponse) VisitPostUrlTitleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type PostUrlTitle500JSONResponse Error

func (response PostUrlTitle500JSONResponse) VisitPostUrlTitleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostUrlUseRequestObject struct {
	Params PostUrlUseParams
	Body   *PostUrlUseJSONRequestBody
}

type PostUrlUseResponseObject interface {
	VisitPostUrlUseResponse(w http.ResponseWriter) error
}

type PostUrlUse200JSONResponse struct {
	Status *int `json:"status,omitempty"`
	Url    *Url `json:"url,omitempty"`
}

func (response PostUrlUse200JSONResponse) VisitPostUrlUseResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUrlUse400JSONResponse Error

func (response PostUrlUse400JSONResponse) VisitPostUrlUseResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostUrlUse401JSONResponse Error

func (response PostUrlUse401JSONResponse) VisitPostUrlUseResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type PostUrlUse500JSONResponse Error

func (response PostUrlUse500JSONResponse) VisitPostUrlUseResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersDescribeRequestObject struct {
	Params GetUsersDescribeParams
}

type GetUsersDescribeResponseObject interface {
	VisitGetUsersDescribeResponse(w http.ResponseWriter) error
}

type GetUsersDescribe200JSONResponse struct {
	Status *int   `json:"status,omitempty"`
	Urls   *[]Url `json:"urls,omitempty"`
	User   *User  `json:"user,omitempty"`
}

func (response GetUsersDescribe200JSONResponse) VisitGetUsersDescribeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersDescribe400JSONResponse Error

func (response GetUsersDescribe400JSONResponse) VisitGetUsersDescribeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersDescribe401JSONResponse Error

func (response GetUsersDescribe401JSONResponse) VisitGetUsersDescribeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersDescribe500JSONResponse Error

func (response GetUsersDescribe500JSONResponse) VisitGetUsersDescribeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersListRequestObject struct {
}

type GetUsersListResponseObject interface {
	VisitGetUsersListResponse(w http.ResponseWriter) error
}

type GetUsersList200JSONResponse struct {
	Status *int    `json:"status,omitempty"`
	Users  *[]User `json:"users,omitempty"`
}

func (response GetUsersList200JSONResponse) VisitGetUsersListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersList400JSONResponse Error

func (response GetUsersList400JSONResponse) VisitGetUsersListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersList401JSONResponse Error

func (response GetUsersList401JSONResponse) VisitGetUsersListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersList500JSONResponse Error

func (response GetUsersList500JSONResponse) VisitGetUsersListResponse(w http.ResponseWriter) error {
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

	// (POST /url/title)
	PostUrlTitle(ctx context.Context, request PostUrlTitleRequestObject) (PostUrlTitleResponseObject, error)

	// (POST /url/use)
	PostUrlUse(ctx context.Context, request PostUrlUseRequestObject) (PostUrlUseResponseObject, error)

	// (GET /users/describe)
	GetUsersDescribe(ctx context.Context, request GetUsersDescribeRequestObject) (GetUsersDescribeResponseObject, error)

	// (GET /users/list)
	GetUsersList(ctx context.Context, request GetUsersListRequestObject) (GetUsersListResponseObject, error)
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

// PostUrlTitle operation middleware
func (sh *strictHandler) PostUrlTitle(ctx echo.Context, params PostUrlTitleParams) error {
	var request PostUrlTitleRequestObject

	request.Params = params

	var body PostUrlTitleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUrlTitle(ctx.Request().Context(), request.(PostUrlTitleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUrlTitle")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUrlTitleResponseObject); ok {
		return validResponse.VisitPostUrlTitleResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUrlUse operation middleware
func (sh *strictHandler) PostUrlUse(ctx echo.Context, params PostUrlUseParams) error {
	var request PostUrlUseRequestObject

	request.Params = params

	var body PostUrlUseJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUrlUse(ctx.Request().Context(), request.(PostUrlUseRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUrlUse")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUrlUseResponseObject); ok {
		return validResponse.VisitPostUrlUseResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsersDescribe operation middleware
func (sh *strictHandler) GetUsersDescribe(ctx echo.Context, params GetUsersDescribeParams) error {
	var request GetUsersDescribeRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsersDescribe(ctx.Request().Context(), request.(GetUsersDescribeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsersDescribe")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersDescribeResponseObject); ok {
		return validResponse.VisitGetUsersDescribeResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsersList operation middleware
func (sh *strictHandler) GetUsersList(ctx echo.Context) error {
	var request GetUsersListRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsersList(ctx.Request().Context(), request.(GetUsersListRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsersList")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersListResponseObject); ok {
		return validResponse.VisitGetUsersListResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
