package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/SarkiMudboy/meeet/internal/models"
	"github.com/SarkiMudboy/meeet/internal/storage"
)

type config interface {
	ServerAddr() string
}

type Application struct {
	config config
	store  *models.Storage
	object storage.ObjectStorage
}

func NewApp(cfg config, store *models.Storage, fileStorage storage.ObjectStorage) *Application {
	return &Application{
		config: cfg,
		store:  store,
		object: fileStorage,
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
