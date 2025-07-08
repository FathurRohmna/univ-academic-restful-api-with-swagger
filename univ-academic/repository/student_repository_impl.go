package repository

import (
	"context"
	"errors"
	"univ-academic/model/domain"

	"gorm.io/gorm"
)

type StudentRepository struct{}

func NewUserRepository() IStudentRepository {
	return &StudentRepository{}
}

func (r *StudentRepository) Save(ctx context.Context, tx *gorm.DB, student domain.Student) (domain.Student, error) {
	err := tx.WithContext(ctx).Save(&student).Error
	if err != nil {
		return domain.Student{}, err
	}

	return student, nil
}

func (r *StudentRepository) FindByEmail(ctx context.Context, tx *gorm.DB, studentEmail string) (domain.Student, error) {
	var student domain.Student
	err := tx.WithContext(ctx).Where("email = ?", studentEmail).First(&student).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Student{}, errors.New("student not found")
		}

		return domain.Student{}, err
	}

	return student, nil
}

func (r *StudentRepository) FindByID(ctx context.Context, tx *gorm.DB, studentID string) (domain.Student, error) {
	var student domain.Student
	err := tx.WithContext(ctx).Where("student_id = ?", studentID).First(&student).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Student{}, errors.New("student not found")
		}
		return domain.Student{}, err
	}
	return student, nil
}
