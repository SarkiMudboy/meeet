package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SarkiMudboy/meeet/database"
	"log"
	"net/http"
)

var ErrUnauthuorized = errors.New("Unauthorized")

func Authorize(r *http.Request) error {

	ctx := context.Background()
	queries := database.New(db)

	session, err := r.Cookie("session_token")
	if err != nil {
		log.Print(err)
		return ErrUnauthuorized
	}

	csrfToken := r.Header.Get("X-CSRF-Token")

	if csrfToken == "" {
		log.Print(err)
		return ErrUnauthuorized
	}
	log.Printf("querying the db with csrf=%s and session token=%s", csrfToken, session)
	arg := database.RetrieveAuthParams{
		CsrfToken:    sql.NullString{String: csrfToken, Valid: true},
		SessionToken: sql.NullString{String: session.Value, Valid: true},
	}
	_, err = queries.RetrieveAuth(ctx, arg)
	if err != nil {
		log.Print(err)
		return ErrUnauthuorized
	}

	return nil
}
