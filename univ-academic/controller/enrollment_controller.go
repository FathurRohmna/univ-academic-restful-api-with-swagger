package controller

import "github.com/labstack/echo/v4"

type IEnrollmentController interface {
	EnrollToCourse(c echo.Context) error
	DeleteEnrollment(c echo.Context) error
}
