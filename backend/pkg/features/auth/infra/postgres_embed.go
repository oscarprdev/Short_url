package postgres

import "embed"

//go:embed sql/queries/select_user_by_id.sql
var selectUserById embed.FS

//go:embed sql/queries/insert_user.sql
var insertUser embed.FS
