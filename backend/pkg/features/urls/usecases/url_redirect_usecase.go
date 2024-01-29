package urls

import (
	"context"
)

func (uc *UrlUsecases) UrlRedirectUsecases(ctx context.Context, url string) (string, error) {
	original, err := uc.Repo.GetOriginalUrl(ctx, url)
	if err != nil {
		return "", err
	}

	return original.OriginalUrl, nil
}
