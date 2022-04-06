package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/wmcff/serve-data/container"
	"github.com/wmcff/serve-data/model/dto"
	"github.com/wmcff/serve-data/service"
)

// ResumeController is a controller for managing resume data.
type ResumeController interface {
	GetResumeList(c echo.Context) error
	NewResume(c echo.Context) error
}

type resumeController struct {
	container container.Container
	service   service.ResumeService
}

// NewResumeController is constructor.
func NewResumeController(container container.Container) ResumeController {
	return &resumeController{container: container, service: service.NewResumeService(container)}
}

// GetResumeList returns the list of all categories.
// @Summary Get a resume list
// @Description Get a resume list
// @Tags Categories
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Resume "Success to fetch a resume list."
// @Failure 401 {string} false "Failed to the authentication."
// @Router /categories [get]
func (controller *resumeController) GetResumeList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllResumes())
}

func (controller *resumeController) NewResume(c echo.Context) error {
	dto := dto.NewResumeDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	fmt.Println(dto.Person)
	fmt.Println(dto.ToString())
	return c.JSON(http.StatusOK, controller.service.FindAllResumes())
}
