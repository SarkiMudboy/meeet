package main

import (
	"log"

	"github.com/SarkiMudboy/meeet/config"
	"github.com/SarkiMudboy/meeet/internal/db"
	"github.com/SarkiMudboy/meeet/internal/handlers"
	"github.com/SarkiMudboy/meeet/internal/models"
	"github.com/SarkiMudboy/meeet/scripts/schema"
)

func main() {

	cfg, err := config.LoadAppConfig()
	if err != nil {
		log.Panic(err)
	}

	db, err := db.New(
		cfg.DB.Addr(),
		cfg.DB.MaxOpenConn(),
		cfg.DB.MaxIdleConn(),
		cfg.DB.MaxConnLifetime(),
	)
	if err != nil {
		log.Panicf("An error occured initializing the database: %s", err.Error())
	}
	defer db.Close()
	log.Println("database connection established")
	schema.RunMigrations()

	store := models.NewStore(db)
	app := handlers.NewApp(&cfg, store)
	router := app.Mount()
	log.Fatal(app.Run(router))
}
