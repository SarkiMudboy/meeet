package schema

import (
	"log"
	"os"

	"github.com/SarkiMudboy/meeet/pkg/env"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {

	db_addr := env.GetString("SCHEMA_DB_ADDR", "mysql://admin:1234@/admin?")

	m, err := migrate.New("file://internal/db/migrations", db_addr)

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		os.Exit(1)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("An error occured: %s", err.Error())
		os.Exit(1)
	}

	log.Println("Migration Sucessful")
}
