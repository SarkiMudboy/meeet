package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SarkiMudboy/meeet/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "meeet:2580@/meeet?parseTime=true")
	if err != nil {
		fmt.Printf("An error occured initializing the database: %s", err.Error())
		os.Exit(1)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// migrate any changes
	runMigrations()
}

func runMigrations() {

	m, err := migrate.New("file://database/migrations", "mysql://meeet:2580@/meeet?")

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
