package user

import (
	"github.com/SarkiMudboy/meeet/models"
	"github.com/SarkiMudboy/meeet/repository"
)

type dbrepository interface {
	createUser(auth models.Auth) error
	getUser(email string) (models.User, error)
	getAuth(email string) (models.Auth, error)
	updateAuth(auth models.Auth) error
}

type Controller struct {
	service dbrepository
}

func InitController(userRepo repository.UserRepository) *Controller {
	return &Controller{
		service: userRepo,
	}
}
