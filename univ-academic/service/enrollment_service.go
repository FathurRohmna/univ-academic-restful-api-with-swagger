package service

import (
	"context"
	"univ-academic/model/web"
)

type IEnrollmentService interface {
	EnrollToCourse(ctx context.Context, studentID string, courseID string) web.EnrollmentResponse
	DeleteEnrollment(ctx context.Context, studentID string, courseID string)
}
