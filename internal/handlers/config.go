package handlers

import (
	"github.com/SarkiMudboy/meeet/internal/models"
	"log"
	"net/http"
	"time"
)

type config interface {
	ServerAddr() string
}

type Application struct {
	config config
	store  *models.Storage
}

func NewApp(cfg config, store *models.Storage) *Application {
	return &Application{
		config: cfg,
		store:  store,
	}
}

func (a *Application) Run(mux http.Handler) error {

	server := &http.Server{
		Addr:         a.config.ServerAddr(),
		Handler:      mux,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 8,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s\n", a.config.ServerAddr())
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
