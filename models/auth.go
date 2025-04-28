package models

type Auth struct {
	AuthId       int16
	Session      string
	PasswordHash string
	CSRFToken    string
}
