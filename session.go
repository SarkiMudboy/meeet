package main

import (
	"errors"
	"net/http"
)

var ErrUnauthuorized = errors.New("Unauthorized")

func Authorize(r *http.Request) error {

	// get the user from db
	// user, err := getUser(req.Email)

	var user *User
	session, err := r.Cookie("session_cookie")
	if err != nil || session.Value == "" || session.Value != user.UserAuth.Session {
		return ErrUnauthuorized
	}

	csrfToken := r.Header.Get("X-CSRF-Token")

	if csrfToken == "" || user.UserAuth.CSRFToken != csrfToken {
		return ErrUnauthuorized
	}

	return nil
}
