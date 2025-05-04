package handlers

import (
	"errors"
	"github.com/SarkiMudboy/meeet/internal/models"
	"net/http"
)

var ErrUnauthuorized = errors.New("Unauthorized")

func (a *Application) Authorize(r *http.Request) (models.Auth, error) {

	session, err := r.Cookie("session_token")
	if err != nil {
		return models.Auth{}, ErrUnauthuorized
	}

	csrfToken := r.Header.Get("X-CSRF-Token")

	if csrfToken == "" {
		return models.Auth{}, ErrUnauthuorized
	}

	auth, err := a.store.Auth.RetrieveAuth(r.Context(), csrfToken, session.Value)
	if err != nil {
		return models.Auth{}, ErrUnauthuorized
	}

	return auth, nil

}
