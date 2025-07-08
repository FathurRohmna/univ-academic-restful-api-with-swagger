package domain

import "time"

type Course struct {
	CourseID     string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"course_id"`
	Name         string    `gorm:"not null" json:"name"`
	Description  string    `gorm:"not null" json:"description"`
	DepartmentID string    `gorm:"type:uuid;not null" json:"department_id"`
	Credits      int       `gorm:"not null" json:"credits"`
	CreatedAt    time.Time `gorm:"default:current_timestamp;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp;not null" json:"updated_at"`

	Enrollments []Enrollment `gorm:"foreignKey:CourseID"`
}
