package controller

import "github.com/labstack/echo/v4"

type ICourseController interface {
	GetAllCourses(c echo.Context) error
}
