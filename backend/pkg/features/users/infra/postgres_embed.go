package postgres

import "embed"

//go:embed sql/queries/select_users.sql
var selectUsers embed.FS

//go:embed sql/queries/select_user_by_id.sql
var selectUserById embed.FS

//go:embed sql/queries/join_user_url.sql
var joinUserUrl embed.FS
