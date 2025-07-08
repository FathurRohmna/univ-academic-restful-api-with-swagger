package repository

import (
	"context"
	"errors"
	"univ-academic/model/domain"
	"univ-academic/model/web"

	"gorm.io/gorm"
)

type CourseRepository struct{}

func NewCourseRepository() ICourseRepository {
	return &CourseRepository{}
}

func (r *CourseRepository) FindCoursesByStudentID(ctx context.Context, tx *gorm.DB, studentID string) ([]web.CourseResponse, error) {
	var courses []web.CourseResponse
	err := tx.WithContext(ctx).
		Table("enrollments").
		Select("courses.course_id, courses.name AS title, enrollments.enrollment_date").
		Joins("JOIN courses ON enrollments.course_id = courses.course_id").
		Where("enrollments.student_id = ?", studentID).
		Scan(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepository) GetAllCourses(ctx context.Context, tx *gorm.DB) ([]web.CourseDetails, error) {
	var courses []web.CourseDetails
	query := `
		SELECT 
			courses.course_id, 
			courses.name AS course_name,
			courses.description AS course_description,
			courses.credits,
			departments.name AS department,
			CONCAT(professors.first_name, ' ', professors.last_name) AS professor_name
		FROM courses
		JOIN departments ON courses.department_id = departments.department_id
		LEFT JOIN teachings ON courses.course_id = teachings.course_id
		LEFT JOIN professors ON teachings.professor_id = professors.professor_id
	`

	err := tx.WithContext(ctx).Raw(query).Scan(&courses).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *CourseRepository) FindByID(ctx context.Context, tx *gorm.DB, courseID string) (domain.Course, error) {
	var course domain.Course
	err := tx.WithContext(ctx).
		Where("course_id = ?", courseID).
		First(&course).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Course{}, errors.New("course does not exist")
		}
		return domain.Course{}, err
	}
	return course, nil
}
