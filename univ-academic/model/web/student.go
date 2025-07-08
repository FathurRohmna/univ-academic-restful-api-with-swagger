package web

import (
	"time"

	"github.com/google/uuid"
)

type CreateStudentRequest struct {
	FirstName string `json:"first_name" validate:"required" example:"John"`
	LastName  string `json:"last_name" validate:"required" example:"Doe"`
	Address   string `json:"address" validate:"required" example:"123 Main St"`
	Email     string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password  string `json:"password" validate:"required,min=8" example:"password123"`
	BirthDate string `json:"birth_date" validate:"required,date" example:"2000-01-01"`
}

type LoginStudentRequest struct {
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"password123"`
}

type StudentResponse struct {
	ID        uuid.UUID `json:"id" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	Address   string    `json:"address" example:"123 Main St"`
	Email     string    `json:"email" example:"john.doe@example.com"`
	BirthDate string    `json:"birth_date" example:"2000-01-01"`
}

type StudentLoginResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Jwt    string    `json:"jwt_tkn"`
}

type StudentWithCoursesResponse struct {
	FullName  string           `json:"full_name"`
	Address   string           `json:"address"`
	BirthDate string           `json:"birth_date"`
	Courses   []CourseResponse `json:"courses"`
}

type CourseResponse struct {
	CourseID       string    `json:"course_id"`
	Title          string    `json:"title"`
	EnrollmentDate time.Time `json:"enrollment_date"`
}
