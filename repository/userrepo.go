package repository

import (
	"context"
	"database/sql"
	"github.com/SarkiMudboy/meeet/database"
	"github.com/SarkiMudboy/meeet/models"
)

type UserRepository struct {
	db *database.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: database.New(db),
	}
}

func (u *UserRepository) createUser(request UserCreateRequest, auth models.Auth) error {

	ctx := context.Background()
	// queries := database.New(db)

	params := database.CreateUserParams{
		Email:    request.Email,
		Password: auth.PasswordHash,
	}

	result, err := u.db.CreateUser(ctx, params)
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
	_, err = u.db.CreateAuth(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) getUser(email string) (models.User, error) {
	ctx := context.Background()
	// queries := database.New(db)

	existing_user := models.User{}
	user, err := u.db.GetUserAuth(ctx, email)
	if err != nil {
		return existing_user, err
	}

	existing_user.User = &database.User{
		UserID: user.UserID,
		Email:  user.Email,
	}
	existing_user.UserAuth = &models.Auth{
		PasswordHash: user.PasswordHash.String,
		CSRFToken:    user.CsrfToken.String,
		Session:      user.SessionToken.String,
	}

	return existing_user, nil
}

func (u *UserRepository) getAuth(email string) (models.Auth, error) {
	ctx := context.Background()
	// queries := database.New(db)

	a := models.Auth{}

	auth, err := u.db.GetAuth(ctx, email)
	if err != nil {
		return a, err
	}

	a.AuthId = auth.AuthID.Int16
	a.PasswordHash = auth.PasswordHash.String
	a.CSRFToken = auth.CsrfToken.String
	a.Session = auth.SessionToken.String

	return a, nil
}

func (u *UserRepository) updateAuth(auth models.Auth) error {

	ctx := context.Background()
	// queries := database.New(db)

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

	_, err := u.db.UpdateUserAuth(ctx, authParams)
	if err != nil {
		return err
	}
	return nil
}
