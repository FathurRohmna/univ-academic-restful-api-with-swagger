package service

import (
	"context"
	"univ-academic/model/web"
)

type IUStudentService interface {
	Login(ctx context.Context, student web.LoginStudentRequest) string
	Register(ctx context.Context, student web.CreateStudentRequest) web.StudentResponse
	GetStudentWithCourses(ctx context.Context, studentID string) (web.StudentWithCoursesResponse, error)
}
