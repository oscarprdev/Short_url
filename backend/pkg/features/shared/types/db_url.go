package types

import (
	"time"

	"github.com/google/uuid"
)

type DbUrl struct {
	Id          uuid.UUID `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	Title       string    `json:"title"`
	Usage       int       `json:"usage"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}
