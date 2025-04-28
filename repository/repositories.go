package repository

import (
	"database/sql"
)

type Repositories struct {
	userRepo *UserRepository
}

func InitRepository(db *sql.DB) *Repositories {
	return &Repositories{
		userRepo: NewUserRepository(db),
	}
}
