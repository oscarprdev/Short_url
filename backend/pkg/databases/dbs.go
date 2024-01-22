package databases

import "database/sql"

type Databases struct {
	SqlDB *sql.DB
}

func ProvideDatabase(db *sql.DB) *Databases {
	return &Databases{
		SqlDB: db,
	}
}
