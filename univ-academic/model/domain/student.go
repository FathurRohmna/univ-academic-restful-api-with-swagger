package domain

import (
	"time"
)

type Student struct {
	StudentID    string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"student_id"`
	FirstName    string    `gorm:"not null" json:"first_name"`
	LastName     string    `gorm:"not null" json:"last_name"`
	Email        string    `gorm:"not null;unique" json:"email"`
	PasswordHash string    `gorm:"not null" json:"password_hash"`
	Address      string    `gorm:"not null" json:"address"`
	DateOfBirth  time.Time `gorm:"not null" json:"date_of_birth"`
	CreatedAt    time.Time `gorm:"default:current_timestamp;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp;not null" json:"updated_at"`

	Enrollments []Enrollment `gorm:"foreignKey:StudentID"`
}
