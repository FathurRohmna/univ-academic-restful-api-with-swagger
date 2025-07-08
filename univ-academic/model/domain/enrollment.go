package domain

import "time"

type Enrollment struct {
	StudentID      string    `gorm:"primaryKey;type:uuid;not null" json:"student_id"`
	CourseID       string    `gorm:"primaryKey;type:uuid;not null" json:"course_id"`
	EnrollmentDate time.Time `gorm:"not null" json:"enrollment_date"`
	CreatedAt      time.Time `gorm:"default:current_timestamp;not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp;not null" json:"updated_at"`

	Student Student `gorm:"foreignKey:StudentID"`
	Course  Course  `gorm:"foreignKey:CourseID"`
}
