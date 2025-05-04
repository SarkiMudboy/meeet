package models

import (
	"context"
	"database/sql"
	"github.com/SarkiMudboy/meeet/internal/models/database"
)

type Auth struct {
	AuthId       int16  `json:"auth_id"`
	Session      string `json:"session"`
	PasswordHash string `json:"passwod_hash"`
	CSRFToken    string `json:"csrf_token"`
}

type AuthRepo interface {
	RetrieveAuth(ctx context.Context, csrfToken, sessionToken string) (Auth, error)
	UpdateAuth(ctx context.Context, auth Auth) error
}

type AuthStore struct {
	db *sql.DB
}

func (a *AuthStore) RetrieveAuth(ctx context.Context, csrfToken, sessionToken string) (Auth, error) {

	queries := database.New(a.db)
	auth := Auth{}
	arg := database.RetrieveAuthParams{
		CsrfToken:    sql.NullString{String: csrfToken, Valid: true},
		SessionToken: sql.NullString{String: sessionToken, Valid: true},
	}
	record, err := queries.RetrieveAuth(ctx, arg)
	if err != nil {
		return auth, err
	}

	auth.AuthId = record.AuthID.Int16
	auth.Session = record.SessionToken.String
	auth.PasswordHash = record.PasswordHash.String
	auth.CSRFToken = record.CsrfToken.String
	return auth, nil

}

func (a *AuthStore) UpdateAuth(ctx context.Context, auth Auth) error {

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
