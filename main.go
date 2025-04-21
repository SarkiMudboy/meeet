package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Auth struct {
	Session      string
	PasswordHash string
	CSRFToken    string
}

type User struct {
	Username string
	Email    string
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

	//check if username exists here

	password, err := HashPassword(req.Password)
	if err != nil {
		httpErr = http.StatusInternalServerError
		http.Error(w, "An error occured", httpErr)
		return
	}

	_ = Auth{PasswordHash: password}
	//persist to user

	fmt.Fprint(w, "Success! New User!")

}
func login(w http.ResponseWriter, r *http.Request) {

	var err error
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

	var user *User // get from db

}
func logout(w http.ResponseWriter, r *http.Request)    {}
func protected(w http.ResponseWriter, r *http.Request) {}

func main() {

	http.HandleFunc("/register", register)
	http.HandleFunc("/login", register)
	http.HandleFunc("/logout", register)
	http.HandleFunc("/protected", register)

	log.Println("Server Listening at :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

}
