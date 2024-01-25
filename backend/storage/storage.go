package storage

import (
	api "short_url/pkg/api"
)

type Storage interface {
	Get(string) *api.User
}
