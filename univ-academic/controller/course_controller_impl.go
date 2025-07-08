package controller

import (
	"net/http"
	"univ-academic/service"

	"github.com/labstack/echo/v4"
)

type CourseController struct {
	CourseService service.ICourseService
}

func NewCourseController(service service.ICourseService) *CourseController {
	return &CourseController{
		CourseService: service,
	}
}

func (controller *CourseController) GetAllCourses(c echo.Context) error {
	courses := controller.CourseService.GetAllCourses(c.Request().Context())
	return c.JSON(http.StatusOK, courses)
}
