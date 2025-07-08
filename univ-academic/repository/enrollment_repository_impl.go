package repository

import (
	"context"
	"errors"
	"time"
	"univ-academic/model/domain"

	"gorm.io/gorm"
)

type EnrollmentRepository struct{}

func NewEnrollmentRepository() IEnrollmentRepository {
	return &EnrollmentRepository{}
}

func (er *EnrollmentRepository) EnrollToCourse(ctx context.Context, tx *gorm.DB, studentID string, courseID string) (domain.Enrollment, error) {
	var existingEnrollment domain.Enrollment
	if err := tx.WithContext(ctx).
		Where("student_id = ? AND course_id = ?", studentID, courseID).
		First(&existingEnrollment).Error; err == nil {
		return domain.Enrollment{}, errors.New("student is already enrolled in this course")
	}

	enrollment := domain.Enrollment{
		StudentID:      studentID,
		CourseID:       courseID,
		EnrollmentDate: time.Now(),
	}

	if err := tx.WithContext(ctx).Create(&enrollment).Error; err != nil {
		return domain.Enrollment{}, err
	}

	return enrollment, nil
}

func (r *EnrollmentRepository) DeleteEnrollment(ctx context.Context, tx *gorm.DB, studentID, courseID string) (domain.Enrollment, error) {
	var enrollment domain.Enrollment
	err := tx.WithContext(ctx).
		Where("student_id = ? AND course_id = ?", studentID, courseID).
		First(&enrollment).Error
	if err != nil {
		return domain.Enrollment{}, errors.New("enrollment not found")
	}

	err = tx.WithContext(ctx).Delete(&enrollment).Error
	if err != nil {
		return domain.Enrollment{}, err
	}

	return enrollment, nil
}
