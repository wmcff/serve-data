package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wmcff/zoogeek/container"
	"github.com/wmcff/zoogeek/model"
	"github.com/wmcff/zoogeek/model/dto"
	"github.com/wmcff/zoogeek/service"
	"github.com/wmcff/zoogeek/session"
)

// UserController is a controller for managing user account.
type UserController interface {
	GetLoginStatus(c echo.Context) error
	GetLoginUser(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type userController struct {
	context      container.Container
	service      service.UserService
	dummyAccount *model.User
}

// NewAccountController is constructor.
func NewUserController(container container.Container) UserController {
	return &userController{
		context:      container,
		service:      service.NewUserService(container),
		dummyAccount: model.NewUserWithPlainPassword("test", "test", 1),
	}
}

// GetLoginStatus returns the status of login.
// @Summary Get the login status.
// @Description Get the login status of current logged-in user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {boolean} bool "The current user have already logged-in. Returns true."
// @Failure 401 {boolean} bool "The current user haven't logged-in yet. Returns false."
// @Router /auth/loginStatus [get]
func (controller *userController) GetLoginStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}

// GetLoginUser returns the account data of logged in user.
// @Summary Get the account data of logged-in user.
// @Description Get the account data of logged-in user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Account "Success to fetch the account data. If the security function is disable, it returns the dummy data."
// @Failure 401 {boolean} bool "The current user haven't logged-in yet. Returns false."
// @Router /auth/loginAccount [get]
func (controller *userController) GetLoginUser(c echo.Context) error {
	if !controller.context.GetConfig().Extension.SecurityEnabled {
		return c.JSON(http.StatusOK, controller.dummyAccount)
	}
	return c.JSON(http.StatusOK, session.GetUser(c))
}

// Login is the method to login using username and password by http post.
// @Summary Login using username and password.
// @Description Login using username and password.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body dto.LoginDto true "User name and Password for logged-in."
// @Success 200 {object} model.Account "Success to the authentication."
// @Failure 401 {boolean} bool "Failed to the authentication."
// @Router /auth/login [post]
func (controller *userController) Login(c echo.Context) error {
	dto := dto.NewLoginDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}

	user := session.GetUser(c)
	if user != nil {
		return c.JSON(http.StatusOK, user)
	}

	authenticate, a := controller.service.AuthenticateByUsernameAndPassword(dto.UserName, dto.Password)
	if authenticate {
		_ = session.SetUser(c, a)
		_ = session.Save(c)
		return c.JSON(http.StatusOK, a)
	}
	return c.NoContent(http.StatusUnauthorized)
}

// Logout is the method to logout by http post.
// @Summary Logout.
// @Description Logout.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /auth/logout [post]
func (controller *userController) Logout(c echo.Context) error {
	_ = session.SetUser(c, nil)
	_ = session.Delete(c)
	return c.NoContent(http.StatusOK)
}
