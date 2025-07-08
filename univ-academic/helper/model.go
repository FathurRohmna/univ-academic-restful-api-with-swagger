package helper

import (
	"log"
	"univ-academic/model/domain"
	"univ-academic/model/web"

	"github.com/google/uuid"
)

func ToStudentResponse(student domain.Student) web.StudentResponse {
	studentID, err := uuid.Parse(student.StudentID)
	if err != nil {
		log.Fatalf("Invalid student ID: %v", err)
	}

	return web.StudentResponse{
		ID:        studentID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
		Address:   student.Address,
		BirthDate: student.DateOfBirth.Format("2006-01-02"),
	}
}

func ToEnrollmentResponse(enrollment domain.Enrollment) web.EnrollmentResponse {
	return web.EnrollmentResponse{
		StudentID:      enrollment.StudentID,
		CourseID:       enrollment.CourseID,
		EnrollmentDate: enrollment.EnrollmentDate,
	}
}
