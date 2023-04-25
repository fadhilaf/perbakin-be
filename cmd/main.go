package main

import (
  "github.com/FadhilAF/perbakin-be/common/env"

  "github.com/FadhilAF/perbakin-be/common/postgres"

  "github.com/FadhilAF/perbakin-be/internal/app"
)

func main() {
	appConfig := env.New(".env")

  postgresDb := postgres.StartPostgresPoolAndMigrate(appConfig.PostgresDSN, "file://config/postgres/migration")

	app := app.New(appConfig, postgresDb)

	app.StartServer()
}
