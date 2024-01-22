package urls

import (
	"crypto/md5"
	"encoding/base64"
	"io"
	"os"
)

func (uc *UrlUsecases) ShortUrlUsecases(ou string) string {
	hasher := md5.New()
	io.WriteString(hasher, ou)
	hash := hasher.Sum(nil)

	encoded := base64.URLEncoding.EncodeToString(hash)

	prodUrl := os.Getenv("PROD_URL")

	shortURL := prodUrl + encoded[:8]

	return shortURL
}
