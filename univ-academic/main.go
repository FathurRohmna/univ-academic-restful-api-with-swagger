package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"univ-academic/app"
	"univ-academic/controller"
	"univ-academic/exception"
	pkgmiddleware "univ-academic/middleware"
	"univ-academic/repository"
	"univ-academic/service"
	"univ-academic/validator"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := app.NewDB()

	courseRepository := repository.NewCourseRepository()

	studentValidator := validator.NewStudentValidator()
	studentRepository := repository.NewUserRepository()
	studentService := service.NewStudentService(studentRepository, courseRepository, db)
	studentController := controller.NewStudentController(studentService, studentValidator)

	enrollmentValidator := validator.NewEnrollmentValidator()
	enrollmentRepository := repository.NewEnrollmentRepository()
	enrollmentService := service.NewEnrollmentService(enrollmentRepository, courseRepository, db)
	enrollmentController := controller.NewEnrollmentController(enrollmentService, enrollmentValidator)

	courseService := service.NewCourseService(courseRepository, db)
	courseController := controller.NewCourseController(courseService)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = exception.ErrorHandler

	api := e.Group("/api")

	studentRouter := api.Group("/students")
	studentRouter.POST("/register", studentController.Register)
	studentRouter.POST("/login", studentController.Login)

	studentRouter.GET("/me", studentController.GetStudentDetails, pkgmiddleware.JWTMiddleware)

	api.POST("/enrollments", enrollmentController.EnrollToCourse, pkgmiddleware.JWTMiddleware)
	api.DELETE("/enrollments/:course_id", enrollmentController.DeleteEnrollment, pkgmiddleware.JWTMiddleware)

	api.GET("/courses", courseController.GetAllCourses, pkgmiddleware.JWTMiddleware)

	go func() {
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Shutting down the server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
	if err := e.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
