package types

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type DbUser struct {
	Id        string              `json:"id"`
	Email     openapi_types.Email `json:"email"`
	Picture   string              `json:"picture"`
	Name      string              `json:"name"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Token     string              `json:"token"`
	ExpiresAt time.Time           `json:"expires_at"`
}
