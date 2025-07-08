package controller

import (
	"net/http"
	"univ-academic/model/web"
	"univ-academic/service"
	"univ-academic/validator"

	"github.com/labstack/echo/v4"
)

type StudentController struct {
	StudentService   service.IUStudentService
	StudentValidator validator.IStudentValidator
}

func NewStudentController(
	studentService service.IUStudentService,
	studentValidator validator.IStudentValidator,
) IStudentController {
	return &StudentController{
		StudentService:   studentService,
		StudentValidator: studentValidator,
	}
}

func (controller *StudentController) Register(c echo.Context) error {
	var studentRequest web.CreateStudentRequest
	if err := c.Bind(&studentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := controller.StudentValidator.ValidateCreateStudent(studentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()
	studentResponse := controller.StudentService.Register(ctx, studentRequest)

	return c.JSON(http.StatusCreated, studentResponse)
}

func (controller *StudentController) Login(c echo.Context) error {
	var loginRequest web.LoginStudentRequest
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := controller.StudentValidator.ValidateLoginStudent(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()
	token := controller.StudentService.Login(ctx, loginRequest)

	return c.JSON(http.StatusOK, map[string]string{"access_token": token})
}

func (c *StudentController) GetStudentDetails(ctx echo.Context) error {
	studentID, ok := ctx.Get("student_id").(string)
	if !ok || studentID == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or missing student ID"})
	}

	response, err := c.StudentService.GetStudentWithCourses(ctx.Request().Context(), studentID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
