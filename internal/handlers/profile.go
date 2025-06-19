package handlers

import (
	"fmt"
	"io"
	"os"

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
		"name":        []string{"min:3", "max:1000", "required"},
		"file:avatar": []string{"ext:jpg,png", "size:10000", "required"},
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

	r.ParseMultipartForm(1024 * 5)
	validationErr := validateAddProfile(r)
	if len(validationErr) > 0 {
		raiseValidationError(w, validationErr)
		return
	}

	displayName := r.PostFormValue("display_name")

	fmt.Println(displayName)

	file, header, err := r.FormFile("avatar")
	if err == nil {
		fmt.Fprintf(w, "%T\n", file)
	}

	defer file.Close()

	diskFile, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, "Server Error", http.StatusBadRequest)
		return
	}

	if _, err := io.Copy(diskFile, file); err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	// get the user somehow and UserID; use authorize maybe?
	// validate that the name has not been taken...(new query?)
	// storage service from a to save the file..log failure here
	// path, err := a.storage.Save(diskFile)
	// a.store.Users.AddProfile(r.Context(), userId int, displayName string, avatarPath string)
	fmt.Fprint(w, "Success! Profile created!")
}
