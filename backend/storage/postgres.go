package storage

import (
	"database/sql"
	"log"
	api "short_url/pkg/api"
)

type PostgresStorage struct {
	SqlDB *sql.DB
}

func NewPostgresStorage(dbUrl string) *PostgresStorage {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can not connect to database", err)
	}

	return &PostgresStorage{
		SqlDB: db,
	}
}

func (s *PostgresStorage) Get(id string) *api.User {
	return &api.User{}
}
