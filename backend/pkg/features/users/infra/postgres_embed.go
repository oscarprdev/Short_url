package postgres

import "embed"

//go:embed sql/queries/select_users.sql
var selectUsers embed.FS
