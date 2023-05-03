package main

import (
	"log"

	"github.com/FadhilAF/perbakin-be/common/env"

	"github.com/FadhilAF/perbakin-be/common/postgres"

	"github.com/FadhilAF/perbakin-be/internal/app"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	appConfig := env.New(".env")

	postgresDb := postgres.StartPostgresPoolAndMigrate(appConfig.PostgresDSN, "config/postgres/migration")

	app := app.New(appConfig, postgresDb)

	app.StartServer()
}
