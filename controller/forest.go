package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wmcff/zoogeek/container"
	"github.com/wmcff/zoogeek/service"
)

// ForestController is a controller for managing forest data.
type ForestController interface {
	GetForestList(c echo.Context) error
}

type forestController struct {
	container container.Container
	service   service.ForestService
}

// NewForestController is constructor.
func NewForestController(container container.Container) ForestController {
	return &forestController{container: container, service: service.NewForestService(container)}
}

// GetForestList returns the list of all categories.
// @Summary Get a forest list
// @Description Get a forest list
// @Tags Categories
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Forest "Success to fetch a forest list."
// @Failure 401 {string} false "Failed to the authentication."
// @Router /categories [get]
func (controller *forestController) GetForestList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllForests())
}
