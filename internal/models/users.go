package models

import (
	"context"
	"database/sql"
	"github.com/SarkiMudboy/meeet/internal/models/database"
	"log"
)

type UserRepo interface {
	CheckUserExists(ctx context.Context, email string) bool
	CreateUser(ctx context.Context, email string, auth Auth) error
	GetUser(ctx context.Context, email string) (User, error)
	GetAuth(ctx context.Context, email string) (Auth, error)
}

type UserStore struct {
	db *sql.DB
}

type User struct {
	User *database.User
	Auth *Auth
}

func (u *UserStore) CheckUserExists(ctx context.Context, email string) bool {

	queries := database.New(u.db)
	r, err := queries.CheckUserExists(ctx, email)
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return false
	}
	return r
}

func (u *UserStore) CreateUser(ctx context.Context, email string, auth Auth) error {

	queries := database.New(u.db)

	params := database.CreateUserParams{
		Email:    email,
		Password: auth.PasswordHash,
	}

	result, err := queries.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	user_id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	arg := database.CreateAuthParams{
		UserID: uint16(user_id),
		PasswordHash: sql.NullString{
			String: auth.PasswordHash,
			Valid:  true,
		},
	}
	_, err = queries.CreateAuth(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

// func CreateUserAuthFromHash() error {
// }

func (u *UserStore) GetUser(ctx context.Context, email string) (User, error) {

	queries := database.New(u.db)
	user := User{}
	userRecord, err := queries.GetUserAuth(ctx, email)
	if err != nil {
		return user, err
	}

	user.User = &database.User{
		UserID: userRecord.UserID,
		Email:  userRecord.Email,
	}
	user.Auth = &Auth{
		PasswordHash: userRecord.PasswordHash.String,
		CSRFToken:    userRecord.CsrfToken.String,
		Session:      userRecord.SessionToken.String,
	}

	return user, nil
}

func (u *UserStore) GetAuth(ctx context.Context, email string) (Auth, error) {
	queries := database.New(u.db)
	a := Auth{}

	auth, err := queries.GetAuth(ctx, email)
	if err != nil {
		return a, err
	}

	a.AuthId = auth.AuthID.Int16
	a.PasswordHash = auth.PasswordHash.String
	a.CSRFToken = auth.CsrfToken.String
	a.Session = auth.SessionToken.String

	return a, nil
}
