package urls

import (
	api "short_url/pkg/api"
	adapters "short_url/pkg/features/urls/adapters"
)

func (uc *UrlUsecases) ShortUrlSimpleUsecases(ou *adapters.OriginalUrl, url *api.Url) error {
	su := shortUrl(ou)

	suDB, err := adapters.AdaptShortUrlToDB(*ou.Ou, su)
	if err != nil {
		return err
	}

	*url = *adapters.AdaptShortUrlToApp(suDB)

	return nil
}
