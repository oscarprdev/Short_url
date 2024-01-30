package seed

import (
	"context"
	"database/sql"
	"fmt"
)

type PostgresRepository struct {
	Db *sql.DB
}

func ProvidePostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		Db: db,
	}
}

var createUrlsTablePath = "sql/create_urls_table.sql"

func (pr *PostgresRepository) createUrlsTable(ctx context.Context) error {
	query, err := createUrlsTable.ReadFile(createUrlsTablePath)
	if err != nil {
		return fmt.Errorf("error reading createUrlsTablePath: %v", err)
	}

	_, err = pr.Db.ExecContext(ctx, string(query))
	if err != nil {
		return fmt.Errorf("error executing create table query: %v", err)
	}

	return nil
}

var createUsersTablePath = "sql/create_users_table.sql"

func (pr *PostgresRepository) createUsersTable(ctx context.Context) error {
	query, err := createUsersTable.ReadFile(createUsersTablePath)
	if err != nil {
		return fmt.Errorf("error reading createUsersTablePath: %v", err)
	}

	_, err = pr.Db.ExecContext(ctx, string(query))
	if err != nil {
		return fmt.Errorf("error executing create table query: %v", err)
	}

	return nil
}

var createUrlUserTablePath = "sql/create_url_user_table.sql"

func (pr *PostgresRepository) createUrlUserTable(ctx context.Context) error {
	query, err := createUrlUserTable.ReadFile(createUrlUserTablePath)
	if err != nil {
		return fmt.Errorf("error reading createUrlUserTablePath: %v", err)
	}

	_, err = pr.Db.ExecContext(ctx, string(query))
	if err != nil {
		return fmt.Errorf("error executing create table query: %v", err)
	}

	return nil
}
