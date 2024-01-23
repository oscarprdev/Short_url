package postgres

import "embed"

//go:embed sql/queries/select_user_by_id.sql
var selectUserById embed.FS

//go:embed sql/queries/insert_user.sql
var insertUser embed.FS

//go:embed sql/queries/update_user.sql
var updateUser embed.FS
