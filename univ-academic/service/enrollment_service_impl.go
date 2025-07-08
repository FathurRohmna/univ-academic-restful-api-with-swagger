package service

import (
	"context"
	"univ-academic/exception"
	"univ-academic/helper"
	"univ-academic/model/web"
	"univ-academic/repository"

	"gorm.io/gorm"
)

type EnrollmentService struct {
	EnrollmentRepository repository.IEnrollmentRepository
	CourseRepository     repository.ICourseRepository
	DB                   *gorm.DB
}

func NewEnrollmentService(enrollmentRepository repository.IEnrollmentRepository, courseRepository repository.ICourseRepository, DB *gorm.DB) *EnrollmentService {
	return &EnrollmentService{
		EnrollmentRepository: enrollmentRepository,
		CourseRepository:     courseRepository,
		DB:                   DB,
	}
}

func (service *EnrollmentService) EnrollToCourse(ctx context.Context, studentID string, courseID string) web.EnrollmentResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err := service.CourseRepository.FindByID(ctx, tx, courseID)
	if err != nil {
		if err.Error() == "course does not exist" {
			panic(exception.NewDataNotFoundError("course does not exist"))
		}

		panic(err)
	}

	enrollment, err := service.EnrollmentRepository.EnrollToCourse(ctx, tx, studentID, courseID)
	if err != nil {
		if err.Error() == "student is already enrolled in this course" {
			panic(exception.NewDataConflictError("student is already enrolled in this course"))
		}


		panic(err)
	}

	return helper.ToEnrollmentResponse(enrollment)
}

func (service *EnrollmentService) DeleteEnrollment(ctx context.Context, studentID string, courseID string) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err := service.EnrollmentRepository.DeleteEnrollment(ctx, tx, studentID, courseID)
	if err != nil {
		if err.Error() == "enrollment not found" {
			panic(exception.NewDataNotFoundError("enrollment not found"))
		}

		panic(err)
	}
}
