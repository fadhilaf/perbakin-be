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

func NewPostgresPool(migrationFilePath, dsn string) *pgxpool.Pool {
  dbpool, err := pgxpool.New(context.Background(), "postgresql://" + dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

  testPostgresConnection(dbpool)

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

func testPostgresConnection(dbpool *pgxpool.Pool) {
	var greeting string
  err := dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
