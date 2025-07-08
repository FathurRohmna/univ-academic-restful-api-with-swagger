package service

import (
	"context"
	"univ-academic/helper"
	"univ-academic/model/web"
	"univ-academic/repository"

	"gorm.io/gorm"
)

type CourseService struct {
	DB               *gorm.DB
	CourseRepository repository.ICourseRepository
}

func NewCourseService(repo repository.ICourseRepository, db *gorm.DB) ICourseService {
	return &CourseService{
		DB:               db,
		CourseRepository: repo,
	}
}

func (service *CourseService) GetAllCourses(ctx context.Context) []web.CourseDetails {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courses, err := service.CourseRepository.GetAllCourses(ctx, tx)
	if err != nil {
		panic(err)
	}

	return courses
}
