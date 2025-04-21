package main

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(h), err
}
