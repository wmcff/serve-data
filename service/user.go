package service

import (
	"github.com/wmcff/zoogeek/container"
	"golang.org/x/crypto/bcrypt"

	"github.com/wmcff/zoogeek/model"
)

type UserService interface {
	AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.User)
}
type userService struct {
	container container.Container
}

func NewUserService(container container.Container) UserService {
	return &userService{container: container}
}

func (a *userService) AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.User) {
	rep := a.container.GetRepository()
	logger := a.container.GetLogger()
	account := model.User{}
	result, err := account.FindByName(rep, username)
	if err != nil {
		logger.GetZapLogger().Errorf(err.Error())
		return false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
		return false, nil
	}

	return true, result
}
