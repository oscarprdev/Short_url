package urls

import (
	api "short_url/pkg/api"
	sharedAdapters "short_url/pkg/features/shared/adapters"
	adapters "short_url/pkg/features/urls/adapters"
)

func (uc *UrlUsecases) ShortUrlSimpleUsecases(ou *adapters.OriginalUrl, url *api.Url, refererUrl string) error {
	su := shortUrl(ou, refererUrl)

	suDB, err := adapters.AdaptShortUrlToDB(*ou.Ou, su)
	if err != nil {
		return err
	}

	*url = *sharedAdapters.AdaptShortUrlToApp(suDB)

	return nil
}
