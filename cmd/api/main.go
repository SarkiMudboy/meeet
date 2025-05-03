package main

import (
	"log"
	"time"

	"github.com/SarkiMudboy/meeet/internal/db"
	"github.com/SarkiMudboy/meeet/internal/env"
	"github.com/SarkiMudboy/meeet/internal/storage"
	"github.com/SarkiMudboy/meeet/scripts/schema"
)

func main() {

	cfg := config{
		addr: env.GetString("SERVER_ADDR", ":8080"),
		db: DBConfig{
			addr:            env.GetString("DB_ADDR", "meeet:2580@/meeet?parseTime=true"),
			maxIdleConn:     env.GetInt("DB_MAX_IDLE_CONN", 10),
			maxOpenConn:     env.GetInt("DB_MAX_OPEN_CONN", 10),
			maxConnLifetime: env.GetInt("DB_MAX_CONN_LIFETIME", 10),
		},
	}

	//db
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConn,
		cfg.db.maxIdleConn,
		time.Duration(cfg.db.maxConnLifetime),
	)
	if err != nil {
		log.Panicf("An error occured initializing the database: %s", err.Error())
	}
	defer db.Close()
	log.Println("database connection established")
	schema.RunMigrations()

	store := storage.NewStore(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	router := app.mount()
	log.Fatal(app.run(router))
}
