package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type DBConfig struct {
	addr            string
	maxIdleConn     int
	maxOpenConn     int
	maxConnLifetime int
}

type application struct {
	config config
}

type config struct {
	addr string
	db   DBConfig
}

func (a *application) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1/auth", func(r chi.Router) {
		r.Post("/register", a.healthCheck)
		r.Post("/login", a.healthCheck)
		r.Post("/logout", a.healthCheck)
		r.Post("/create-meeting", a.healthCheck)
	})

	r.Get("/health", a.healthCheck)
	return r
}

func (a *application) run(mux http.Handler) error {

	server := &http.Server{
		Addr:         a.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 8,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s\n", a.config.addr)
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
