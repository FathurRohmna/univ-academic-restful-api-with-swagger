package controller

import (
	"net/http"
	"univ-academic/model/web"
	"univ-academic/service"
	"univ-academic/validator"

	"github.com/labstack/echo/v4"
)

type EnrollmentController struct {
	EnrollmentService   service.IEnrollmentService
	EnrollmentValidator validator.IEnrollmentValidator
}

func NewEnrollmentController(
	enrollmentService service.IEnrollmentService, enrollmentValidator validator.IEnrollmentValidator,
) IEnrollmentController {
	return &EnrollmentController{
		EnrollmentService:   enrollmentService,
		EnrollmentValidator: enrollmentValidator,
	}
}

func (controller *EnrollmentController) EnrollToCourse(c echo.Context) error {
	studentID, ok := c.Get("student_id").(string)
	if !ok || studentID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or missing student ID"})
	}

	var request web.EnrollToCourseRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := controller.EnrollmentValidator.ValidateEnrollToCourse(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()
	enrollment := controller.EnrollmentService.EnrollToCourse(ctx, studentID, request.CourseID)

	return c.JSON(http.StatusCreated, enrollment)
}

func (controller *EnrollmentController) DeleteEnrollment(c echo.Context) error {
	studentID, ok := c.Get("student_id").(string)
	courseID := c.Param("course_id")

	if !ok || studentID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or missing student ID"})
	}

	ctx := c.Request().Context()
	controller.EnrollmentService.DeleteEnrollment(ctx, studentID, courseID)

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
