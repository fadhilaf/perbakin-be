package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func StartPostgresPoolAndMigrate(migrationFilePath, dsn string) *pgxpool.Pool {
  dbpool, err := pgxpool.New(context.Background(), "postgresql://" + dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

  migratePostgres(migrationFilePath, dsn)

  return dbpool
}

func migratePostgres(migrationFilePath, dsn string) {
  migration, err := migrate.New(migrationFilePath, "pgx://" + dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln(err)
	}
}
