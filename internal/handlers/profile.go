package handlers

import (
	"fmt"
	"os"
	"path"

	// "github.com/SarkiMudboy/meeet/internal/models"
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
	// "github.com/SarkiMudboy/meeet/internal/utils"
	// "github.com/thedevsaddam/govalidator"
)

type ProfileRequest struct {
	Name   string   `json:"name"`
	Avatar *os.File `json:"avatar"`
}

func validateAddProfile(r *http.Request) url.Values {

	rules := govalidator.MapData{
		"display_name": []string{"min:3", "max:1000", "required"},
		"file:avatar":  []string{"ext:jpg,png", "size:10000", "required"},
	}

	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		RequiredDefault: true,
	}
	validator := govalidator.New(opts)
	e := validator.Validate()

	return e
}

func (a *Application) addProfileInformation(w http.ResponseWriter, r *http.Request) {

	var httpErr int
	if r.Method != http.MethodPost {
		httpErr = http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", httpErr)
		return
	}

	userAuth, err := a.Authorize(r)
	if err != nil {
		httpErr = http.StatusUnauthorized
		http.Error(w, "Unauthorized", httpErr)
		return
	}

	r.ParseMultipartForm(1024 * 5)
	validationErr := validateAddProfile(r)
	if len(validationErr) > 0 {
		raiseValidationError(w, validationErr)
		return
	}

	displayName := r.PostFormValue("display_name")

	//name not been taken
	if a.store.Users.CheckUserExists(r.Context(), "", displayName) {
		httpErr = http.StatusConflict
		http.Error(w, "A user with that name already exists", httpErr)
		return
	}

	// get avatar
	file, header, err := r.FormFile("avatar")
	if err != nil {
		fmt.Fprintf(w, "%T\n", file)
	}

	defer file.Close()

	// create dest file
	diskFile, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Save to disk
	fileName := fmt.Sprintf("%s%s", displayName, path.Ext(header.Filename))
	path, err := a.object.Save(fileName, "avatars/", diskFile)

	if err != nil {
		httpErr = http.StatusInternalServerError
		http.Error(w, "Server error", httpErr)
		return
	}

	err = a.store.Users.AddProfile(r.Context(), int(userAuth.UserId), displayName, path)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest) // wrap error here
		return
	}
	fmt.Fprint(w, "Success! Profile created!")
}
