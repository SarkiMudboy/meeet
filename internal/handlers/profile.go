package handlers

import (
	"fmt"
	// "github.com/SarkiMudboy/meeet/internal/models"
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
	// "github.com/SarkiMudboy/meeet/internal/utils"
	// "github.com/thedevsaddam/govalidator"
)

type ProfileRequest struct {
	Name   string
	Avatar []byte
}

func validateAddProfile(r *http.Request) (url.Values, *ProfileRequest) {

	var request ProfileRequest
	rules := govalidator.MapData{
		"name":        []string{"min:3", "max:1000", "required"},
		"file:avatar": []string{"ext:jpg,png", "size:10000", "required"},
	}

	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		RequiredDefault: true,
		// Data:            &request,
	}
	validator := govalidator.New(opts)
	e := validator.Validate()

	return e, &request
}

func (a *Application) addProfileInformation(w http.ResponseWriter, r *http.Request) {

	var httpErr int
	if r.Method != http.MethodPost {
		httpErr = http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", httpErr)
		return
	}

	validationErr, profile := validateAddProfile(r)
	if len(validationErr) > 0 {
		raiseValidationError(w, validationErr)
		return
	}
	fmt.Printf("%v", profile)

	// may have to still do this
	// a.store.Users.AddProfile(r.Context(), userId int, displayName string, avatarPath string)
	// r.ParseMultipartForm(1024 * 5)
	// data := r.MultipartForm
	// displayName := r.PostFormValue("display_name")
	// fmt.Println(displayName)
	// file, _, err := r.FormFile("avatar")
	// if err == nil {
	// 	fmt.Fprintf(w, "%v", file)
	// }

	// get the user somehow and UserID; use authorize maybe?
	// use a goroutine via the storage service from a to save the file..log failure here

	http.Error(w, "Invalid Method", httpErr)
	// return
}
