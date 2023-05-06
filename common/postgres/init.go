package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func StartPostgresPoolAndMigrate(dsn, migrationFilePath string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), "postgresql://"+dsn)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	migratePostgres(migrationFilePath, dsn)

	return dbpool
}

func migratePostgres(migrationFilePath, dsn string) {
	migration, err := migrate.New("file://"+migrationFilePath, "pgx://"+dsn)
	if err != nil {
		log.Fatalln("Error ketika membuat koneksi database untuk migrasi:", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln("Error ketika melakukan migrasi:", err)
	}
}
