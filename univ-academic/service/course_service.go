package service

import (
	"context"
	"univ-academic/model/web"
)

type ICourseService interface {
	GetAllCourses(ctx context.Context) []web.CourseDetails
}
