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

type LoginRequest struct {
	Email    string
	Password string
}

func register(w http.ResponseWriter, r *http.Request) {

	var err error
	var httpErr int

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

	//check if user exists here
	if exists := checkUserExists(req.Email); exists {
		httpErr = http.StatusBadRequest
		http.Error(w, "A user with that email already exists", httpErr)
		return
	}

	password, err := HashPassword(req.Password)
	if err != nil {
		httpErr = http.StatusInternalServerError
		http.Error(w, "An error occured", httpErr)
		return
	}

	auth := Auth{PasswordHash: password}

	//persist to user
	createUser(req, auth)
	fmt.Fprint(w, "Success! New User!")

}
func login(w http.ResponseWriter, r *http.Request) {

	// var err error
	var httpErr int
	var req LoginRequest

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

	user, err := getUser(req.Email)
	if err != nil {
		httpErr = http.StatusBadRequest
		http.Error(w, "Invalid login credentials", httpErr)
		return
	}

	if !CheckPassword([]byte(user.UserAuth.PasswordHash), req.Password) {
		var httpErr = http.StatusBadRequest
		http.Error(w, "Invalid Username/Password", httpErr)
		return
	}

	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

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
	auth, err := getAuth(req.Email)
	if err != nil {
		httpErr := http.StatusInternalServerError
		http.Error(w, "An error occured", httpErr)
		return
	}

	auth.Session = sessionToken
	auth.CSRFToken = csrfToken

	if err = updateAuth(auth); err != nil {
		httpErr := http.StatusInternalServerError
		http.Error(w, "An error occured", httpErr)
		return
	}

	fmt.Fprint(w, "Login Success!")
}

func logout(w http.ResponseWriter, r *http.Request) {

	var httpErr int

	if err := Authorize(r); err != nil {
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

	// clear auth data from user record
	fmt.Fprint(w, "Log out successful")

}

func createMeeting(w http.ResponseWriter, r *http.Request) {

	var httpErr int

	if r.Method != http.MethodPost {
		httpErr = http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", httpErr)
		return
	}

	if err := Authorize(r); err != nil {
		httpErr = http.StatusUnauthorized
		http.Error(w, "Invalid request", httpErr)
		return
	}

	fmt.Fprint(w, "Create meeting success")
}

func main() {

	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/create-meeting", createMeeting)

	log.Println("Server Listening at :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

}
