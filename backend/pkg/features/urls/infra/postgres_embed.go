package postgres

import "embed"

//go:embed sql/schemas/url_schema.sql
var urlSchema embed.FS

//go:embed sql/schemas/url_user_schema.sql
var urlUserSchema embed.FS

//go:embed sql/schemas/url_statistics_schema.sql
var urlStatisticsSchema embed.FS

//go:embed sql/queries/insert_url.sql
var insertUrlQuery embed.FS

//go:embed sql/queries/select_user_by_id.sql
var selectUserById embed.FS

//go:embed sql/queries/insert_url_user.sql
var insertUrlUserQuery embed.FS
