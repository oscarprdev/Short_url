package adapters

import (
	"time"

	"github.com/google/uuid"
)

type ShortUrlDBBody struct {
	Id          uuid.UUID `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type OriginalUrl struct {
	Ou *string
}

type AuthUserInfo struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
