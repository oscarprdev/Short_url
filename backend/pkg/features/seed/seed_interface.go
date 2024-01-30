package seed

import "context"

type SeedRepository interface {
	createUrlsTable(ctx context.Context) error
	createUsersTable(ctx context.Context) error
	createUrlUserTable(ctx context.Context) error
}
