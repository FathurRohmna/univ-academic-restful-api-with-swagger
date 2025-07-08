package validator

import (
	"errors"
	"time"
	"univ-academic/model/web"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IStudentValidator interface {
	ValidateCreateStudent(req web.CreateStudentRequest) error
	ValidateLoginStudent(req web.LoginStudentRequest) error
}

type studentValidator struct{}

func NewStudentValidator() IStudentValidator {
	return &studentValidator{}
}

func isValidDate(value interface{}) error {
	str, _ := value.(string)
	_, err := time.Parse("2006-01-02", str)
	if err != nil {
		return errors.New("birth date must be in YYYY-MM-DD format")
	}
	return nil
}

func (sv *studentValidator) ValidateCreateStudent(req web.CreateStudentRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.FirstName,
			validation.Required.Error("first name is required"),
			validation.RuneLength(1, 50).Error("first name must be between 1 and 50 characters"),
		),
		validation.Field(
			&req.LastName,
			validation.Required.Error("last name is required"),
			validation.RuneLength(1, 50).Error("last name must be between 1 and 50 characters"),
		),
		validation.Field(
			&req.Address,
			validation.Required.Error("address is required"),
		),
		validation.Field(
			&req.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email format"),
		),
		validation.Field(
			&req.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("password must be between 8 and 30 characters"),
		),
		validation.Field(
			&req.BirthDate,
			validation.Required.Error("birth date is required"),
			validation.By(isValidDate),
		),
	)
}

func (sv *studentValidator) ValidateLoginStudent(req web.LoginStudentRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email format"),
		),
		validation.Field(
			&req.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("password must be between 8 and 30 characters"),
		),
	)
}
