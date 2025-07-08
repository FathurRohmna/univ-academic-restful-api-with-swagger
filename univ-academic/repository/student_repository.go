package repository

import (
	"context"
	"univ-academic/model/domain"

	"gorm.io/gorm"
)

type IStudentRepository interface {
	Save(ctx context.Context, tx *gorm.DB, student domain.Student) (domain.Student, error)
	FindByEmail(ctx context.Context, tx *gorm.DB, studentEmail string) (domain.Student, error)
	FindByID(ctx context.Context, tx *gorm.DB, studentID string) (domain.Student, error)
}
