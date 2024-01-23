package adapters

import (
	"fmt"
	"regexp"
	api "short_url/pkg/api"
	errors "short_url/pkg/features/shared/handlers"
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

func validateUrl(ou string) (*OriginalUrl, error) {
	urlPattern := `^https?://([a-zA-Z0-9.-]+)(:[0-9]+)?/?([^\s]*)$`

	regex, err := regexp.Compile(urlPattern)
	if err != nil {
		return nil, &errors.BadRequestError{
			Details: fmt.Sprintf("Original url is not a valid url: %v", err),
		}
	}

	match := regex.MatchString(ou)
	if !match {
		return nil, &errors.BadRequestError{
			Details: fmt.Sprintf("Original url is not a valid url: %v", err),
		}
	}

	return &OriginalUrl{
		Ou: &ou,
	}, nil
}

func AdaptShortUrlToDB(ou, su string) (*ShortUrlDBBody, error) {
	validOu, err := validateUrl(ou)
	if err != nil {
		return nil, err
	}

	title := "Default url title"
	now := time.Now().UTC()
	uuid := uuid.New()

	return &ShortUrlDBBody{
		Id:          uuid,
		OriginalUrl: *validOu.Ou,
		ShortUrl:    su,
		CreatedAt:   now,
		UpdatedAt:   now,
		Title:       title,
		ExpiresAt:   now.Add(7 * 24 * time.Hour), // 1 week
	}, nil
}

func AdaptShortUrlToApp(subody *ShortUrlDBBody) *api.Url {
	return &api.Url{
		OriginalUrl: &subody.OriginalUrl,
		ShortUrl:    &subody.ShortUrl,
		UpdatedAt:   &subody.UpdatedAt,
		CreatedAt:   &subody.CreatedAt,
		ExpiresAt:   &subody.ExpiresAt,
		TitleUrl:    &subody.Title,
	}
}
