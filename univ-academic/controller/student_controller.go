package controller

import "github.com/labstack/echo/v4"

type IStudentController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	GetStudentDetails(ctx echo.Context) error
}
