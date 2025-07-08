package repository

import (
	"context"
	"univ-academic/model/domain"
	"univ-academic/model/web"

	"gorm.io/gorm"
)

type ICourseRepository interface {
	FindCoursesByStudentID(ctx context.Context, tx *gorm.DB, studentID string) ([]web.CourseResponse, error)
	GetAllCourses(ctx context.Context, tx *gorm.DB) ([]web.CourseDetails, error)
	FindByID(ctx context.Context, tx *gorm.DB, courseID string) (domain.Course, error)
}
