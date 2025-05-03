package storage

import (
	"context"
	"database/sql"
	"github.com/SarkiMudboy/meeet/internal/storage/database"
)

type Auth struct {
	AuthId       int16  `json:"auth_id"`
	Session      string `json:"session"`
	PasswordHash string `json:"passwod_hash"`
	CSRFToken    string `json:"csrf_token"`
}

type AuthRepo interface {
	updateAuth(ctx context.Context, auth Auth) error
}

type AuthStore struct {
	db *sql.DB
}

func (a *AuthStore) updateAuth(ctx context.Context, auth Auth) error {

	queries := database.New(a.db)

	authParams := database.UpdateUserAuthParams{
		AuthID: sql.NullInt16{Int16: auth.AuthId, Valid: true},
		SessionToken: sql.NullString{
			String: auth.Session,
			Valid:  true,
		},
		CsrfToken: sql.NullString{
			String: auth.CSRFToken,
			Valid:  true,
		},
		PasswordHash: sql.NullString{
			String: auth.PasswordHash,
			Valid:  true,
		},
	}

	_, err := queries.UpdateUserAuth(ctx, authParams)
	if err != nil {
		return err
	}
	return nil
}
