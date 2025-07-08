package web

import "time"

type EnrollmentResponse struct {
	StudentID      string    `json:"student_id"`
	CourseID       string    `json:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
}

type EnrollToCourseRequest struct {
	CourseID string `json:"course_id" validate:"required,uuid" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
}
