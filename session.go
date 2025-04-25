package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SarkiMudboy/meeet/database"
	"net/http"
)

var ErrUnauthuorized = errors.New("Unauthorized")

func Authorize(r *http.Request) (Auth, error) {

	ctx := context.Background()
	queries := database.New(db)
	auth := Auth{}

	session, err := r.Cookie("session_token")
	if err != nil {
		return Auth{}, ErrUnauthuorized
	}

	csrfToken := r.Header.Get("X-CSRF-Token")

	if csrfToken == "" {
		return Auth{}, ErrUnauthuorized
	}

	arg := database.RetrieveAuthParams{
		CsrfToken:    sql.NullString{String: csrfToken, Valid: true},
		SessionToken: sql.NullString{String: session.Value, Valid: true},
	}
	a, err := queries.RetrieveAuth(ctx, arg)
	if err != nil {
		return Auth{}, ErrUnauthuorized
	}

	auth.AuthId = a.AuthID.Int16
	auth.Session = a.SessionToken.String
	auth.PasswordHash = a.PasswordHash.String
	auth.CSRFToken = a.CsrfToken.String
	return auth, nil

}
