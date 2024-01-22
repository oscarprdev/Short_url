package adapters

import (
	"fmt"
	"regexp"
	errors "short_url/pkg/features/shared/handlers"

	"github.com/google/uuid"
)

type ShortUrlDBBody struct {
	Id          uuid.UUID `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
}

func validateUrl(ou string) error {
	urlPattern := `^https?://([a-zA-Z0-9.-]+)(:[0-9]+)?/?([^\s]*)$`

	regex, err := regexp.Compile(urlPattern)
	if err != nil {
		return &errors.BadRequestError{
			Details: fmt.Sprintf("Original url is not a valid url: %v", err),
		}
	}

	match := regex.MatchString(ou)
	if !match {
		return &errors.BadRequestError{
			Details: fmt.Sprintf("Original url is not a valid url: %v", err),
		}
	}

	return nil
}

func AdaptShortUrlToDB(ou, su string) (*ShortUrlDBBody, error) {
	err := validateUrl(ou)
	if err != nil {
		return nil, err
	}

	uuid := uuid.New()
	return &ShortUrlDBBody{
		Id:          uuid,
		OriginalUrl: ou,
		ShortUrl:    su,
	}, nil
}
