package main

import (
	"encoding/json"
	"fmt"
	"github.com/SarkiMudboy/meeet/database"
	"log"
	"net/http"
	"time"
)

type Auth struct {
	AuthId       int16
	Session      string
	PasswordHash string
	CSRFToken    string
}

type User struct {
	User     *database.User
	UserAuth *Auth
}

type UserCreateRequest struct {
	Email    string
	Password string
}
