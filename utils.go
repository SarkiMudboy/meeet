package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(h), err
}

func CheckPassword(PasswordHash []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(PasswordHash, []byte(password))
	if err != nil {
		return false
	}

	return true
}

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("An error occured: %s", err.Error())
	}

	return base64.URLEncoding.EncodeToString(bytes)
}
