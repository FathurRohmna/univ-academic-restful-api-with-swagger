package repository

import (
	"context"
	"univ-academic/model/domain"

	"gorm.io/gorm"
)

type IEnrollmentRepository interface {
	EnrollToCourse(ctx context.Context, tx *gorm.DB, studentID string, courseID string) (domain.Enrollment, error)
	DeleteEnrollment(ctx context.Context, tx *gorm.DB, studentID string, courseID string) (domain.Enrollment, error)
}
