package seed

import "embed"

//go:embed sql/create_urls_table.sql
var createUrlsTable embed.FS

//go:embed sql/create_users_table.sql
var createUsersTable embed.FS

//go:embed sql/create_url_user_table.sql
var createUrlUserTable embed.FS
