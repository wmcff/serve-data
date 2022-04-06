package router

import (
	"net/http"

	"github.com/wmcff/serve-data/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wmcff/serve-data/container"
)

func Init(e *echo.Echo, container container.Container) {
	setCORSConfig(e, container)
	setErrorController(e, container)
	setResumeController(e, container)
	setUserController(e, container)
}

func setCORSConfig(e *echo.Echo, container container.Container) {
	if container.GetConfig().Extension.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     []string{"*"},
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}))
	}
}

func setErrorController(e *echo.Echo, container container.Container) {
	errorHandler := controller.NewErrorController(container)
	e.HTTPErrorHandler = errorHandler.JSONError
	e.Use(middleware.Recover())
}

func setUserController(e *echo.Echo, container container.Container) {
	user := controller.NewUserController(container)
	e.GET(controller.APIUserLoginStatus, func(c echo.Context) error { return user.GetLoginStatus(c) })
	e.GET(controller.APIUserLoginUser, func(c echo.Context) error { return user.GetLoginUser(c) })

	if container.GetConfig().Extension.SecurityEnabled {
		e.POST(controller.APIUserLogin, func(c echo.Context) error { return user.Login(c) })
		e.POST(controller.APIUserLogout, func(c echo.Context) error { return user.Logout(c) })
	}
}

func setResumeController(e *echo.Echo, container container.Container) {
	category := controller.NewResumeController(container)
	e.GET(controller.APIResumes, func(c echo.Context) error { return category.GetResumeList(c) })
}
