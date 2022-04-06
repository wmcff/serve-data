package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wmcff/serve-data/container"
)

type APIError struct {
	Code    int
	Message string
}

type ErrorController interface {
	JSONError(err error, c echo.Context)
}

type errorController struct {
	container container.Container
}

func NewErrorController(container container.Container) ErrorController {
	return &errorController{container: container}
}

func (controller *errorController) JSONError(err error, c echo.Context) {
	logger := controller.container.GetLogger()
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}
	var apierr APIError
	apierr.Code = code
	apierr.Message = msg

	if !c.Response().Committed {
		if reserr := c.JSON(code, apierr); reserr != nil {
			logger.GetZapLogger().Errorf(reserr.Error())
		}
	}
	logger.GetZapLogger().Debugf(err.Error())
}
