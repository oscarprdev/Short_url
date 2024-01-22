package postgres

import "database/sql"

type PostgresRepository struct {
	Db *sql.DB
}

func ProvidePostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		Db: db,
	}
}
