package service

import (
	"context"
	"os"
	"time"
	"univ-academic/exception"
	"univ-academic/helper"
	"univ-academic/model/domain"
	"univ-academic/model/web"
	"univ-academic/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StudentService struct {
	StudentRepository repository.IStudentRepository
	CourseRepository  repository.ICourseRepository
	DB                *gorm.DB
}

func NewStudentService(studentRepository repository.IStudentRepository, courseRepository repository.ICourseRepository, DB *gorm.DB) *StudentService {
	return &StudentService{
		StudentRepository: studentRepository,
		CourseRepository:  courseRepository,
		DB:                DB,
	}
}

func (service *StudentService) Login(ctx context.Context, studentRequest web.LoginStudentRequest) string {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FindByEmail(ctx, tx, studentRequest.Email)
	if err != nil {
		if err.Error() == "student not found" {
			panic(exception.NewInvalidCredentialError("Email or password is not valid"))
		}

		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(student.PasswordHash), []byte(studentRequest.Password))
	if err != nil {
		panic(exception.NewInvalidCredentialError("Email or password is not valid"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"student_id": student.StudentID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func (service *StudentService) Register(ctx context.Context, studentRequest web.CreateStudentRequest) web.StudentResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err := service.StudentRepository.FindByEmail(ctx, tx, studentRequest.Email)
	if err == nil {
		panic(exception.NewInvalidCredentialError("email is already registered"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(studentRequest.Password), 10)
	if err != nil {
		panic(err)
	}

	birthDate, err := time.Parse("2006-01-02", studentRequest.BirthDate)
	if err != nil {
		panic(err)
	}

	newStudent := domain.Student{
		Email:        studentRequest.Email,
		PasswordHash: string(hash),
		FirstName:    studentRequest.FirstName,
		LastName:     studentRequest.LastName,
		Address:      studentRequest.Address,
		DateOfBirth:  birthDate,
	}

	newStudent, err = service.StudentRepository.Save(ctx, tx, newStudent)
	if err != nil {
		panic(err)
	}

	return helper.ToStudentResponse(newStudent)
}

func (s *StudentService) GetStudentWithCourses(ctx context.Context, studentID string) (web.StudentWithCoursesResponse, error) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	student, err := s.StudentRepository.FindByID(ctx, tx, studentID)
	if err != nil {
		return web.StudentWithCoursesResponse{}, err
	}

	courses, err := s.CourseRepository.FindCoursesByStudentID(ctx, tx, studentID)
	if err != nil {
		return web.StudentWithCoursesResponse{}, err
	}

	var courseResponses []web.CourseResponse
	for _, course := range courses {
		courseResponses = append(courseResponses, web.CourseResponse{
			CourseID:       course.CourseID,
			Title:          course.Title,
			EnrollmentDate: course.EnrollmentDate,
		})
	}

	return web.StudentWithCoursesResponse{
		FullName:  student.FirstName + " " + student.LastName,
		Address:   student.Address,
		BirthDate: student.DateOfBirth.Format("2006-01-02"),
		Courses:   courseResponses,
	}, nil
}
