package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/SarkiMudboy/meeet/internal/models"
	"github.com/SarkiMudboy/meeet/internal/utils"
	"github.com/thedevsaddam/govalidator"
)

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string
	Password string
}

func raiseValidationError(w http.ResponseWriter, e url.Values) {

	err := map[string]interface{}{"ValidationErr": e}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)

}

func validateRegister(r *http.Request) url.Values {

	var requestType UserCreateRequest

	rules := govalidator.MapData{
		"email":    []string{"required", "min:5", "email"},
		"password": []string{"required", "min:3", "max:50"},
	}
	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		RequiredDefault: true,
		Data:            &requestType,
	}
	validator := govalidator.New(opts)
	e := validator.ValidateJSON()

	return e
}

func (a *Application) register(w http.ResponseWriter, r *http.Request) {

	var err error
	var httpErr int
	ctx := r.Context()

	if r.Method != http.MethodPost {
		httpErr = http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", httpErr)
		return
	}

	var req UserCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpErr = http.StatusInternalServerError
		http.Error(w, "An error occured", httpErr)
		return
	}

	// validate here
	validationError := validateRegister(r)
	if len(validationError) > 0 {
		raiseValidationError(w, validationError)
		return
	}

	//check if user exists here
	if exists := a.store.Users.CheckUserExists(ctx, req.Email); exists {
		httpErr = http.StatusBadRequest
		http.Error(w, "A user with that email already exists", httpErr)
		return
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		httpErr = http.StatusInternalServerError
		http.Error(w, "An error occured", httpErr)
		return
	}

	auth := models.Auth{PasswordHash: password}

	//persist to user
	a.store.Users.CreateUser(ctx, req.Email, auth)
	fmt.Fprint(w, "Success! New User!")

}
func (a *Application) login(w http.ResponseWriter, r *http.Request) {

	// var err error
	var httpErr int
	var req LoginRequest
	ctx := r.Context()

	if r.Method != http.MethodPost {
		httpErr = http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", httpErr)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpErr = http.StatusInternalServerError
		http.Error(w, "An error occured", httpErr)
		return
	}

	user, err := a.store.Users.GetUser(ctx, req.Email)
	if err != nil {
		httpErr = http.StatusBadRequest
		http.Error(w, "Invalid Username/Password", httpErr)
		return
	}

	if !utils.CheckPassword([]byte(user.Auth.PasswordHash), req.Password) {
		var httpErr = http.StatusBadRequest
		http.Error(w, "Invalid Username/Password", httpErr)
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(time.Duration(time.Hour * 24)),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(time.Duration(time.Hour * 24)),
		HttpOnly: false,
	})

	//Save this session and token to user Auth
	auth, err := a.store.Users.GetAuth(ctx, req.Email)
	if err != nil {
		httpErr := http.StatusInternalServerError
		http.Error(w, "An error occured", httpErr)
		return
	}

	auth.Session = sessionToken
	auth.CSRFToken = csrfToken

	if err = a.store.Auth.UpdateAuth(ctx, auth); err != nil {
		httpErr := http.StatusInternalServerError
		http.Error(w, "An error occured", httpErr)
		return
	}

	fmt.Fprint(w, "Login Success!")
}

func (a *Application) logout(w http.ResponseWriter, r *http.Request) {

	var httpErr int
	var userAuth models.Auth
	ctx := r.Context()

	userAuth, err := a.Authorize(r)
	if err != nil {
		httpErr = http.StatusUnauthorized
		http.Error(w, "Invalid request", httpErr)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	userAuth.Session = ""
	userAuth.CSRFToken = ""
	if err = a.store.Auth.UpdateAuth(ctx, userAuth); err != nil {
		httpErr = http.StatusBadRequest
		http.Error(w, "Invalid request", httpErr)
		return
	}

	fmt.Fprint(w, "Log out successful")

}

func (a *Application) createMeeting(w http.ResponseWriter, r *http.Request) {

	var httpErr int

	if r.Method != http.MethodPost {
		httpErr = http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", httpErr)
		return
	}

	if _, err := a.Authorize(r); err != nil {
		httpErr = http.StatusUnauthorized
		http.Error(w, "Invalid request", httpErr)
		return
	}

	fmt.Fprint(w, "Create meeting success")
}
