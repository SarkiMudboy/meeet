package models

import "github.com/SarkiMudboy/meeet/database"

type User struct {
	User     *database.User
	UserAuth *Auth
}
