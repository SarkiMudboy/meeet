package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SarkiMudboy/meeet/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "meeet:2580@/meeet?parseTime=true")
	if err != nil {
		fmt.Printf("An error occured initializing the database: %s", err.Error())
		os.Exit(1)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// migrate any changes
	runMigrations()
}

func runMigrations() {

	m, err := migrate.New("file://database/migrations", "mysql://meeet:2580@/meeet?")

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		os.Exit(1)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("An error occured: %s", err.Error())
		os.Exit(1)
	}

	log.Println("Migration Sucessful")
}

func checkUserExists(email string) bool {
	ctx := context.Background()
	queries := database.New(db)

	r, err := queries.CheckUserExists(ctx, email)
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return false
	}
	return r
}

func createUser(request UserCreateRequest, auth Auth) error {

	ctx := context.Background()
	queries := database.New(db)

	params := database.CreateUserParams{
		Email:    request.Email,
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

func getUser(email string) (User, error) {
	ctx := context.Background()
	queries := database.New(db)

	u := User{}
	user, err := queries.GetUserAuth(ctx, email)

	if err != nil {
		return u, err
	}

	u.User = &database.User{
		UserID: user.UserID,
		Email:  user.Email,
	}
	u.UserAuth = &Auth{
		PasswordHash: user.PasswordHash.String,
		CSRFToken:    user.CsrfToken.String,
		Session:      user.SessionToken.String,
	}

	return u, nil
}

func getAuth(email string) (Auth, error) {
	ctx := context.Background()
	queries := database.New(db)

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

func updateAuth(auth Auth) error {

	ctx := context.Background()
	queries := database.New(db)

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
