package urls

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	adapters "short_url/pkg/features/urls/adapters"
	"time"
)

type OriginalUrl struct {
	Ou *string
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func hashString(input string) string {
	hasher := md5.New()

	io.WriteString(hasher, input)

	hash := hasher.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

func shortUrl(ou *adapters.OriginalUrl, refererUrl string) string {
	// Compute the MD5 hash of the original URL
	hasher := md5.New()
	io.WriteString(hasher, *ou.Ou)
	hash := hasher.Sum(nil)

	// Encode the hash to a base64 string
	encoded := base64.URLEncoding.EncodeToString(hash)

	// Combine with a timestamp and a random string for more complexity
	timestamp := time.Now().UnixNano()
	randomStr := generateRandomString(4)

	// Combine the components to form the short URL
	shortURL := fmt.Sprintf("%s/%s%s-%d", refererUrl, randomStr, encoded[:8], timestamp)

	// Hash the concatenated string and take the first 8 characters
	finalHash := hashString(shortURL)
	shortURL = fmt.Sprintf("%s%s", refererUrl, finalHash[:8])

	return shortURL
}
