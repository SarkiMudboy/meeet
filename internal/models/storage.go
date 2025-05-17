package models

import "database/sql"

type Storage struct {
	Users UserRepo
	Auth  AuthRepo
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		Users: &UserStore{db},
		Auth:  &AuthStore{db},
	}
}
