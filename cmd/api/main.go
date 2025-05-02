package main

import (
	"github.com/SarkiMudboy/meeet/internal/env"
	"log"
)

func main() {

	cfg := config{
		addr: env.GetEnv("SERVER_ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}
	router := app.mount()

	log.Fatal(app.run(router))
}
